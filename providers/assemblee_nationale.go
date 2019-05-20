package providers

import (
	"strconv"
	"strings"

	"cran/domain"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type assemblee_nationale struct{}

func (p *assemblee_nationale) Name() string {
	return "Assemblée Nationale Française"
}

func (p *assemblee_nationale) Accept(URL string) bool {
	return strings.Contains(URL, "assemblee-nationale.fr")
}

func parseNode(report *domain.Report, queue *[]*domain.Section, ele *goquery.Selection) {
	var parent domain.Node

	if len(*queue) == 0 {
		parent = report
	} else {
		parent = (*queue)[len(*queue)-1]
	}

	if ele.HasClass("ouverture_seance") ||
		ele.HasClass("Point") ||
		ele.HasClass("intervention") {
		title := ele.ChildrenFiltered("h5.presidence, h2").First()

		if title.Text() != "" {
			lvl := 1

			if titleClass, ok := title.Attr("class"); ok {
				if l, err := strconv.Atoi(titleClass[5:]); err == nil {
					lvl = l
				}
			}

			section := domain.NewSection(strings.TrimSpace(title.Text()), lvl)

			lastSection, _ := parent.(*domain.Section)

			if lastSection != nil && section.Level <= lastSection.Level {
				idx := 0
				for i, p := range *queue {
					if section.Level <= p.Level {
						break
					}

					idx = i
					parent = p
				}

				if idx == 0 {
					parent = report
				}

				*queue = append((*queue)[:idx], (*queue)[idx+1:]...)
			}

			parent.Append(section)

			*queue = append(*queue, section)
		}

		ele.Children().Each(func(_ int, c *goquery.Selection) {
			parseNode(report, queue, c)
		})
	}

	if ele.Is("p") {
		content, _ := ele.Html()
		parent.Append(domain.NewIntervention("", content))
	}
}

func (p *assemblee_nationale) Fetch(URL string, callback domain.ProviderCallback) {
	report := domain.NewReport(URL, URL)
	queue := make([]*domain.Section, 0)

	c := colly.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		report.Title = e.Text
	})

	c.OnHTML(".SYCERON > .ouverture_seance, .SYCERON > .Point", func(ele *colly.HTMLElement) {
		parseNode(report, &queue, ele.DOM)
	})

	c.OnScraped(func(r *colly.Response) {
		// for _, s := range queue {
		// 	report.Append(s)
		// }
		callback(report, nil)
	})

	c.Visit(URL)
}
