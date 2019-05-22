package assembleenationale

import (
	"cran/domain"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type parser struct {
	index     int
	collector *colly.Collector
	report    *domain.Report
	queue     []domain.Node
}

func newParser(url string) *parser {
	r := domain.NewReport(url, url, "http://www.assemblee-nationale.fr/commun/ceresian/images/logo-an.png")

	return &parser{
		report:    r,
		collector: colly.NewCollector(),
		queue:     []domain.Node{r},
	}
}

func (p *parser) nextIndex(prefix string) string {
	p.index++

	return fmt.Sprintf("%s%d", prefix, p.index)
}

func (p *parser) parseNode(ele *goquery.Selection) {
	if ele.HasClass("ouverture_seance") ||
		ele.HasClass("Point") ||
		ele.HasClass("intervention") {

		// Ok, that's a container, let's find if it is the start of a new section!

		title := ele.ChildrenFiltered("h5.presidence, h2").First()

		if title.Text() != "" {
			level := 1

			// Try to extract the current nested level since class follow this structure
			// - titre1
			// - titre2
			// - titre 99
			if titleClass, ok := title.Attr("class"); ok {
				if lvl, err := strconv.Atoi(titleClass[5:]); err == nil {
					level = lvl
				}
			}

			section := domain.NewSection(p.nextIndex("section-"), strings.TrimSpace(title.Text()), level)

			// Find the first element in queue with a level inferior to ours
			var i int

			for i = len(p.queue) - 1; i >= 0; i-- {
				node := p.queue[i]

				// We've reached the report, that's the root node so exits now!
				if node == p.report {
					break
				}

				if s, isSection := node.(*domain.Section); isSection && s.Level < section.Level {
					break
				}
			}

			// And cut the queue till this element
			p.queue = append(p.queue[:i+1], section)

			// Append this new section to its parent
			p.queue[len(p.queue)-2].Append(section)
		}

		ele.Children().Each(func(_ int, c *goquery.Selection) {
			p.parseNode(c)
		})
	} else if ele.Is("p") {
		author := ele.Find("a[href]")

		// And remove the node so it doesn't appear in the final sentence
		author.Remove()
		authorName := author.Text()
		content, _ := ele.Html()
		content = strings.TrimSpace(content)

		// Remove noisy stuff which is anything before the first period
		if idx := strings.Index(content, "."); idx > -1 && authorName != "" {
			content = strings.TrimSpace(content[idx+1:])
		}

		parent := p.queue[len(p.queue)-1]

		if authorName != "" {
			parent.Append(domain.NewIntervention(p.nextIndex("intervention-"), authorName, content))

			d := p.collector.Clone()

			d.OnHTML("html", func(profileEle *colly.HTMLElement) {
				title := profileEle.DOM.Find("h1")
				img := profileEle.DOM.Find(".deputes-image img")

				p.report.AddSpeaker(authorName,
					title.Text(),
					profileEle.Request.URL.String(),
					profileEle.Request.AbsoluteURL(img.AttrOr("src", "")),
					strings.TrimSpace(title.Next().Text()),
					strings.TrimSpace(img.Parent().Parent().Children().Last().Text()))
			})

			if authorLink, exists := author.Attr("href"); exists {
				d.Visit(authorLink)
			}
		} else {
			parent.Append(domain.NewNotice(p.nextIndex("notice-"), content))
		}
	}
}

func (p *parser) run(cb domain.ProviderCallback) {
	p.collector.OnHTML("h1", func(e *colly.HTMLElement) {
		p.report.Title = e.Text
	})

	p.collector.OnHTML(".SYCERON > .ouverture_seance, .SYCERON > .Point", func(ele *colly.HTMLElement) {
		p.parseNode(ele.DOM)
	})

	p.collector.OnScraped(func(r *colly.Response) {
		cb(p.report, nil)
	})

	p.collector.Visit(p.report.URL)
}
