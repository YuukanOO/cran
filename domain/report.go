package domain

// Speaker hold profile information about a people.
type Speaker struct {
	ID         string
	Name       string
	ProfileURL string
	PictureURL string
	Location   string
	Side       string
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

// AddSpeaker adds a speaker to this report and returns the instantiated object.
func (r *Report) AddSpeaker(id, name, profileURL, pictureURL, location, side string) *Speaker {
	speaker := &Speaker{
		ID:         id,
		Name:       name,
		ProfileURL: profileURL,
		PictureURL: pictureURL,
		Location:   location,
		Side:       side,
	}

	r.Speakers[id] = speaker

	return speaker
}
