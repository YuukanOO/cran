package generating

import "errors"

// DoneCallback represents the callback being called when a parsing has ended.
type DoneCallback func(report *Report, err error)

// Provider represents a report provider which is able to returns an instance
// of a report from an URL by parsing the page content.
type Provider interface {
	Accept(url string) bool
	Name() string
	Fetch(url string, callback DoneCallback)
}

var (
	// I know this global registry is arguable but I think it makes perfect sense
	// in this tiny domain.
	providers = make([]Provider, 0)

	// ErrProviderNotFound when no provider has been found to process the request
	ErrProviderNotFound = errors.New("Provider not found")
)

// Register provider implementations to the inner registry to make it available.
func Register(instances ...Provider) {
	providers = append(providers, instances...)
}

// GuessProvider try to find a provider which is able to process an URL.
func GuessProvider(url string) (Provider, error) {
	for _, p := range providers {
		if p.Accept(url) {
			return p, nil
		}
	}

	return nil, ErrProviderNotFound
}

// GetProvider retrieves a provider by its name.
func GetProvider(name string) (Provider, error) {
	for _, p := range providers {
		if p.Name() == name {
			return p, nil
		}
	}

	return nil, ErrProviderNotFound
}
