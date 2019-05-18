package domain

import "errors"

// ProviderCallback represents the callback being called when a parsing has ended.
type ProviderCallback func(report *Report, err error)

// Provider represents a report provider which is able to returns an instance
// of report from an URL by parsing the page content.
type Provider interface {
	Accept(URL string) bool
	Name() string
	Fetch(URL string, callback ProviderCallback)
}

var (
	providers = make([]Provider, 0)

	// ErrProviderNotFound when no provider has been found to process the request
	ErrProviderNotFound = errors.New("Provider not found")
)

// Register provider implementations to the registry.
func Register(instances ...Provider) {
	providers = append(providers, instances...)
}

// GuessProvider try to find a provider which is able to process an URL.
func GuessProvider(URL string) (Provider, error) {
	for _, p := range providers {
		if p.Accept(URL) {
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
