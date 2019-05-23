package domain

import "fmt"

// Speaker hold profile information about a people.
type Speaker struct {
	ID         string
	Name       string
	ProfileURL string
	PictureURL string
	Location   string
	Side       string
}

// Meta retrieves some meta information about the speaker.
func (s *Speaker) Meta() string {
	return fmt.Sprintf("%s · %s", s.Side, s.Location)
}

// Report represents the root node of the tree and hold informations
// about the currently being generated report such as speakers.
type Report struct {
	Title    string
	URL      string
	IconURL  string
	Speakers map[string]*Speaker
	children []Node
}

// NewReport instantiates a new report.
func NewReport(title, URL, iconURL string) *Report {
	return &Report{
		Title:    title,
		URL:      URL,
		IconURL:  iconURL,
		Speakers: make(map[string]*Speaker),
		children: make([]Node, 0),
	}
}

func (r *Report) Append(nodes ...Node) { r.children = append(r.children, nodes...) }
func (r *Report) Children() []Node     { return r.children }
func (*Report) Type() string           { return "Report" }
func (*Report) ID() string             { return "root" }

// AddSpeaker adds a speaker to this report.
func (r *Report) AddSpeaker(speaker *Speaker) {
	r.Speakers[speaker.ID] = speaker
}

// Speaker retrieves a speaker by its ID.
func (r *Report) Speaker(name string) *Speaker {
	return r.Speakers[name]
}
