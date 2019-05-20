package domain

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeProvider struct {
	name   string
	accept string
}

func (p *fakeProvider) Accept(URL string) bool                      { return strings.Contains(URL, p.accept) }
func (p *fakeProvider) Name() string                                { return p.name }
func (p *fakeProvider) Fetch(URL string, callback ProviderCallback) {}

var (
	p1 = &fakeProvider{name: "p1", accept: "www.a.domain.fr"}
	p2 = &fakeProvider{name: "p2", accept: "www.assemblee-nationale.fr"}
	p3 = &fakeProvider{name: "p3", accept: "www.something.else.io"}
)

func TestMain(m *testing.M) {
	Register(p1, p2)

	r := m.Run()

	providers = make([]Provider, 0)

	os.Exit(r)
}

func TestRegisterProvider(t *testing.T) {
	assert.Len(t, providers, 2, "It should contains 2 providers")

	Register(p3)

	assert.Len(t, providers, 3, "It should contains 3 providers")
}

func TestGetProvider(t *testing.T) {
	p, _ := GetProvider(p1.name)

	assert.Equal(t, p1, p)

	p, _ = GetProvider(p2.name)

	assert.Equal(t, p2, p)
}

func TestGuessProvider(t *testing.T) {
	p, _ := GuessProvider("http://www.assemblee-nationale.fr/15/cri/2018-2019/20190229.asp")

	assert.Equal(t, p2, p)
}
