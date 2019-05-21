// Package assembleenationale provides a parser for reports made by the french assemblée nationale.
package assembleenationale

import (
	"strings"

	"cran/domain"
)

type assembleeNationaleProvider struct{}

func (p *assembleeNationaleProvider) Name() string {
	return "Assemblée Nationale Française"
}

func (p *assembleeNationaleProvider) Accept(URL string) bool {
	return strings.Contains(URL, "assemblee-nationale.fr")
}

func (p *assembleeNationaleProvider) Fetch(URL string, callback domain.ProviderCallback) {
	parser := newParser(URL)
	parser.run(callback)
}

func init() {
	domain.Register(&assembleeNationaleProvider{})
}
