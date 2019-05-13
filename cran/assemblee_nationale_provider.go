package cran

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type assembleeNationale struct{}

func (p *assembleeNationale) Accept(URL string) bool {
	return strings.Contains(URL, "assemblee-nationale.fr")
}

func (p *assembleeNationale) Name() string {
	return "Assemblée Nationale"
}

func (p *assembleeNationale) Fetch(URL string, callback ProviderCallback) {
	report := newReport(URL)

	c := colly.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		report.Title = e.Text
	})

	c.OnHTML(".Point", func(ele *colly.HTMLElement) {
		sectionTitle := ele.DOM.Find("h2").First().Text()

		// Some sections appears to be empty and it mess everything
		if sectionTitle == "" {
			return
		}

		section := report.addSection(sectionTitle)

		ele.DOM.Find("p").Each(func(i int, e *goquery.Selection) {
			d := c.Clone()

			// Find the author
			author := e.Find("a[href]")

			// And remove the node so it doesn't appear in the final sentence
			author.Remove()

			authorName := author.Text()

			// Append the intervention
			sentence, _ := e.Html()

			// Remove noisy stuff which is anything before the first period
			if idx := strings.Index(sentence, "."); idx > -1 && authorName != "" {
				sentence = strings.TrimLeft(sentence[idx+1:], " ")
			}

			section.addIntervention(authorName, sentence)

			// Upon profile visit success, keep some information on it
			d.OnHTML("html", func(profileEle *colly.HTMLElement) {
				title := profileEle.DOM.Find("h1")
				img := profileEle.DOM.Find(".deputes-image img")

				report.addSpeaker(authorName,
					title.Text(),
					profileEle.Request.URL.String(),
					profileEle.Request.AbsoluteURL(img.AttrOr("src", "")),
					title.Next().Text(),
					img.Parent().Parent().Last().Text())
			})

			// And visits its profile
			if authorLink, exists := author.Attr("href"); exists {
				d.Visit(authorLink)
			}
		})
	})

	c.OnScraped(func(r *colly.Response) {
		callback(report, nil)
	})

	c.Visit(URL)
}
