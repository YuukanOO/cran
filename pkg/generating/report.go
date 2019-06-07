package generating

// Node represents a parent -> children hierarchy and is implemented
// by all elements being generated.
type Node interface {
	ID() string
	Type() string
	Append(...Node)
	Children() []Node
}

// ReportType represents the report node type.
const ReportType = "Report"

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
func NewReport(url string) *Report {
	return &Report{
		URL:      url,
		Speakers: make(map[string]*Speaker),
		children: make([]Node, 0),
	}
}

// AddSpeakers adds one or many speaker to this report.
func (r *Report) AddSpeakers(speakers ...*Speaker) {
	for _, s := range speakers {
		r.Speakers[s.ID] = s
	}
}

func (*Report) ID() string             { return "root" }
func (r *Report) Type() string         { return ReportType }
func (r *Report) Append(nodes ...Node) { r.children = append(r.children, nodes...) }
func (r *Report) Children() []Node     { return r.children }
