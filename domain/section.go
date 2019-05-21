package domain

// Section holds data about a specific portion of the report.
type Section struct {
	Title    string
	Level    int
	children []Node
}

// NewSection instantiates a new section with the given data.
func NewSection(title string, level int) *Section {
	return &Section{
		Title: title,
		Level: level,
	}
}

func (s *Section) Append(nodes ...Node) { s.children = append(s.children, nodes...) }
func (s *Section) Children() []Node     { return s.children }
func (*Section) Type() string           { return "Section" }
