package generating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const reportURL = "https://some.report.url"

func TestNewReport(t *testing.T) {
	r := NewReport(reportURL)

	assert.Equal(t, "root", r.ID())
	assert.Equal(t, ReportType, r.Type())
}

func TestAddSpeakerToReport(t *testing.T) {
	r := NewReport(reportURL)
	assert.Equal(t, reportURL, r.URL)

	s := NewSpeaker("1", "https://some.profile.url")

	r.AddSpeakers(s)
	assert.Len(t, r.Speakers, 1)
	assert.Equal(t, s, r.Speakers[s.ID])
}

func TestAddChildrenToReport(t *testing.T) {
	r := NewReport(reportURL)
	s1 := NewSection("1", "A section", 1)
	s2 := NewSection("2", "Another one", 1)

	r.Append(s1, s2)

	assert.Len(t, r.Children(), 2)
	assert.Equal(t, s1, r.Children()[0])
	assert.Equal(t, s2, r.Children()[1])
}
