package cran

import "errors"

// ProviderCallback represents the callback being called when a parsing has ended.
type ProviderCallback func(report *Report, err error)

// ReportProvider defines an interface to fetch report from an URL.
type ReportProvider interface {
	Accept(URL string) bool
	Name() string
	Fetch(URL string, callback ProviderCallback)
}

var providers = []ReportProvider{
	&assembleeNationale{},
}

// ErrProviderNotFound is thrown when a provider for the given name can not be found
var ErrProviderNotFound = errors.New("Provider not found")

// GuessProvider tries to find a report provider for the given URL.
func GuessProvider(URL string) (ReportProvider, error) {
	for _, p := range providers {
		if p.Accept(URL) {
			return p, nil
		}
	}

	return nil, ErrProviderNotFound
}

// Provider retrieve a provider by its name.
func Provider(name string) (ReportProvider, error) {
	for _, p := range providers {
		if p.Name() == name {
			return p, nil
		}
	}

	return nil, ErrProviderNotFound
}
