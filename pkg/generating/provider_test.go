package generating

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockProvider struct {
	name   string
	accept string
	result *Report
}

func (p *mockProvider) Accept(url string) bool {
	return strings.Contains(url, p.accept)
}

func (p *mockProvider) Name() string {
	return p.name
}

func (p *mockProvider) Fetch(url string, callback DoneCallback) {
	if p.result != nil {
		callback(p.result, nil)
	} else {
		callback(nil, errCouldNotGenerate)
	}
}

const (
	pErrURL    = "www.a.domain.fr/reporturl"
	pResultURL = "www.assemblee-nationale.fr/reporturl"
)

var (
	errCouldNotGenerate = errors.New("Could not generate the report")
	pErr                = &mockProvider{name: "p1", accept: "www.a.domain.fr"}
	pResult             = &mockProvider{
		name:   "p2",
		accept: "www.assemblee-nationale.fr",
		result: NewReport(pResultURL),
	}
	p3 = &mockProvider{name: "p3", accept: "www.something.else.io"}
)

func TestMain(m *testing.M) {
	Register(pErr, pResult)

	r := m.Run()

	providers = make([]Provider, 0)

	os.Exit(r)
}

func TestRegisterProvider(t *testing.T) {
	assert.Len(t, providers, 2, "It should contains 2 providers before registration")

	Register(p3)

	assert.Len(t, providers, 3, "It should contains 3 providers after registration")
}

func TestGetProvider(t *testing.T) {
	p, _ := GetProvider(pErr.name)

	assert.Equal(t, pErr, p, "It should get the first provider")

	p, _ = GetProvider(pResult.name)

	assert.Equal(t, pResult, p, "It should get the second provider")

	_, err := GetProvider("does not exist")

	if assert.NotNil(t, err) {
		assert.Equal(t, ErrProviderNotFound, err)
	}
}

func TestGuessProvider(t *testing.T) {
	p, _ := GuessProvider("http://www.assemblee-nationale.fr/15/cri/2018-2019/20190229.asp")

	assert.Equal(t, pResult, p, "It should get the second provider by guessing it")

	_, err := GuessProvider("http://unknown.url/something")

	if assert.NotNil(t, err) {
		assert.IsType(t, ErrProviderNotFound, err)
	}
}
