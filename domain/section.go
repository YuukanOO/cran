package domain

// Section holds data about a specific portion of the report.
type Section struct {
	Title    string
	children []Node
}

// NewSection instantiates a new section with the given data.
func NewSection(title string) *Section {
	return &Section{
		Title: title,
	}
}

func (s *Section) Append(nodes ...Node) { s.children = append(s.children, nodes...) }
func (s *Section) Children() []Node     { return s.children }
