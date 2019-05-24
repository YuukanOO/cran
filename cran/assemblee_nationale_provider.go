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
	speakers := make(map[string]*Speaker)
	report := newReport(URL)

	c := colly.NewCollector(colly.Async(true))

	c.OnHTML(".SYCERON h1", func(e *colly.HTMLElement) {
		report.Title = e.Text
	})

	c.OnHTML("#deputes-fiche", func(e *colly.HTMLElement) {
		speaker := speakers[e.Request.URL.String()]

		title := e.DOM.Find("h1")
		img := e.DOM.Find(".deputes-image img")

		speaker.Name = title.Text()
		speaker.URL = e.Request.URL.String()
		speaker.PictureURL = e.Request.AbsoluteURL(img.AttrOr("src", ""))
		speaker.Location = strings.TrimSpace(title.Next().Text())
		speaker.Side = strings.TrimSpace(img.Parent().Parent().Children().Last().Text())
	})

	c.OnHTML(".Point", func(ele *colly.HTMLElement) {
		sectionTitle := ele.DOM.Find("h2").First().Text()

		// Some sections appears to be empty and it mess everything
		if sectionTitle == "" {
			return
		}

		section := report.addSection(sectionTitle)

		ele.DOM.Find("p").Each(func(i int, e *goquery.Selection) {
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

			if authorLink, authorLinkExists := author.Attr("href"); authorLinkExists {
				if _, speakerExists := speakers[authorLink]; !speakerExists {
					speakers[authorLink] = &Speaker{
						ID: authorName,
					}

					ele.Request.Visit(authorLink)
				}
			}
		})
	})

	c.Visit(URL)
	c.Wait()

	for _, speaker := range speakers {
		report.Speakers[speaker.ID] = speaker
	}

	callback(report, nil)
}
