package cran

type assembleNationaleCRProvider struct{}

func (p *assembleNationaleCRProvider) Fetch(URL string, callback ProviderCallback) {
	callback(nil, nil)
}