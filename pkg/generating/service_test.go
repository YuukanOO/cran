package generating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGenerateWithNoProvider(t *testing.T) {
	s := NewService()

	s.Generate("https://an.unknown.site/report", func(r *Report, err error) {
		assert.Nil(t, r, "Report should be nil")

		if assert.NotNil(t, err, "It should have an error") {
			assert.Equal(t, ErrProviderNotFound, err, "It should be an ErrProviderNotFound error")
		}
	})
}

func TestServiceGenerateWithError(t *testing.T) {
	s := NewService()

	s.Generate(pErrURL, func(r *Report, err error) {
		assert.Nil(t, r, "Report should be nil")

		if assert.NotNil(t, err, "It should have an error") {
			assert.Equal(t, errCouldNotGenerate, err, "It should be a generation error")
		}
	})
}

func TestServiceGenerateWithSuccess(t *testing.T) {
	s := NewService()

	s.Generate(pResultURL, func(r *Report, err error) {
		assert.Nil(t, err, "It should not have any errors")
		assert.NotNil(t, r, "It should have a valid report")
	})
}
