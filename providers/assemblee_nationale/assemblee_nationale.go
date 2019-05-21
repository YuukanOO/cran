// Package assembleenationale provides a parser for reports made by the french assemblée nationale.
package assembleenationale

import (
	"fmt"
	"strconv"
	"strings"

	"cran/domain"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

type assembleeNationaleProvider struct{}

func (p *assembleeNationaleProvider) Name() string {
	return "Assemblée Nationale Française"
}

func (p *assembleeNationaleProvider) Accept(URL string) bool {
	return strings.Contains(URL, "assemblee-nationale.fr")
}

func parseNode(report *domain.Report, index func() int, queue []domain.Node, ele *goquery.Selection) []domain.Node {
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

			section := domain.NewSection(fmt.Sprintf("s%d", index()), strings.TrimSpace(title.Text()), level)

			// Find the first element in queue with a level inferior to ours
			var i int

			for i = len(queue) - 1; i >= 0; i-- {
				node := queue[i]

				// We've reached the report, that's the root node so exits now!
				if node == report {
					break
				}

				if s, isSection := node.(*domain.Section); isSection && s.Level < section.Level {
					break
				}
			}

			// And cut the queue till this element
			queue = append(queue[:i+1], section)

			// Append this new section to its parent
			queue[len(queue)-2].Append(section)
		}

		ele.Children().Each(func(_ int, c *goquery.Selection) {
			queue = parseNode(report, index, queue, c)
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

		parent := queue[len(queue)-1]

		if authorName != "" {
			parent.Append(domain.NewIntervention(fmt.Sprintf("i%d", index()), authorName, content))
		} else {
			parent.Append(domain.NewNotice(fmt.Sprintf("n%d", index()), content))
		}
	}

	return queue
}

func (p *assembleeNationaleProvider) Fetch(URL string, callback domain.ProviderCallback) {
	report := domain.NewReport(URL, URL, "http://www.assemblee-nationale.fr/commun/ceresian/images/logo-an.png")

	// Since the HTML is a mess, this queue holds sections
	queue := []domain.Node{report}
	
	index := 0
	generator := func () int {
		index++
		return index
	}

	c := colly.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		report.Title = e.Text
	})

	c.OnHTML(".SYCERON > .ouverture_seance, .SYCERON > .Point", func(ele *colly.HTMLElement) {
		queue = parseNode(report, generator, queue, ele.DOM)
	})

	c.OnScraped(func(r *colly.Response) {
		callback(report, nil)
	})

	c.Visit(URL)
}

func init() {
	domain.Register(&assembleeNationaleProvider{})
}
