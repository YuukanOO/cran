package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	reportName          = "Report Name"
	reportURL           = "https://some.domain/report"
	iconURL             = "https://some.domain/icon.png"
	section1Title       = "Section 1"
	section2Title       = "Section 2"
	interventionSpeaker = "John Doe"
	interventionContent = "This is me!"
	noticeContent       = "(closed)"
)

func TestNodeManipulations(t *testing.T) {
	assert := assert.New(t)
	report := NewReport(reportName, reportURL, iconURL)
	section1 := NewSection("s1", section1Title, 1)
	section2 := NewSection("s2", section2Title, 1)
	intervention1 := NewIntervention("i1", interventionSpeaker, interventionContent)
	intervention2 := NewIntervention("i2", interventionSpeaker, interventionContent)
	notice := NewNotice("n1", noticeContent)

	// Append intervention to the second section
	section2.Append(intervention1)
	// Append the second section and an intervention to the first one
	section1.Append(section2, intervention2)
	// And finally adds the section 1 to the report
	report.Append(section1, notice)

	reportChildren := report.Children()

	if assert.Len(reportChildren, 2, "Report should have 2 children") {
		section, _ := reportChildren[0].(*Section)

		if assert.NotNil(section) {
			assert.Equal(section1Title, section.Title, "Title should match")

			children := section.Children()

			if assert.Len(children, 2, "Section 1 should have 2 children") {
				section, _ = children[0].(*Section)

				if assert.NotNil(section, "It should have one section") {
					assert.Equal(section2Title, section.Title, "Title should match")

					subChildren := section.Children()

					if assert.Len(subChildren, 1, "Section 2 should have one intervention") {
						intervention, _ := subChildren[0].(*Intervention)

						if assert.NotNil(intervention) {
							assert.Equal(interventionSpeaker, intervention.SpeakerID, "Speaker should match")
							assert.Equal(interventionContent, intervention.Content, "Content should match")
						}
					}
				}

				intervention, _ := children[1].(*Intervention)

				if assert.NotNil(intervention, "It should have one intervention") {
					assert.Equal(interventionSpeaker, intervention.SpeakerID, "Speaker should match")
					assert.Equal(interventionContent, intervention.Content, "Content should match")
				}
			}
		}

		notice, _ := reportChildren[1].(*Notice)

		if assert.NotNil(notice) {
			assert.Equal(noticeContent, notice.Content)
		}
	}
}
