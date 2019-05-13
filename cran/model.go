package cran

import "fmt"

// Report is the main structure representing what have been parsed by a provider.
type Report struct {
	Title    string
	Source   string
	Speakers map[string]*Speaker
	Sections []*Section
}

// Section represents a specific section of the report.
type Section struct {
	ID            string
	Title         string
	Interventions []*Intervention
}

// Intervention is a specific speech intervention made by a speaker.
type Intervention struct {
	ID        string
	SpeakerID string
	Sentence  string
}

// Speaker represents well... a speaker!
type Speaker struct {
	ID         string
	Name       string
	URL        string
	PictureURL string
	Location   string
	Side       string
}

func newReport(source string) *Report {
	return &Report{
		Source:   source,
		Speakers: make(map[string]*Speaker),
		Sections: make([]*Section, 0),
	}
}

func (r *Report) addSpeaker(ID, name, URL, pictureURL, location, side string) {
	r.Speakers[ID] = &Speaker{
		ID:         ID,
		Name:       name,
		URL:        URL,
		PictureURL: pictureURL,
		Location:   location,
		Side:       side,
	}
}

func (r *Report) addSection(title string) *Section {
	section := &Section{
		ID:            fmt.Sprintf("section-%d", len(r.Sections)),
		Title:         title,
		Interventions: make([]*Intervention, 0),
	}

	r.Sections = append(r.Sections, section)

	return section
}

func (s *Section) addIntervention(speaker, sentence string) {
	s.Interventions = append(s.Interventions, &Intervention{
		ID:        fmt.Sprintf("%s--intervention-%d", s.ID, len(s.Interventions)),
		SpeakerID: speaker,
		Sentence:  sentence,
	})
}

// Speaker retrieves a speaker by its ID.
func (r *Report) Speaker(ID string) *Speaker {
	return r.Speakers[ID]
}

// Meta retrieves some meta information about the speaker.
func (s *Speaker) Meta() string {
	return fmt.Sprintf("%s · %s", s.Side, s.Location)
}
