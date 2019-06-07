package generating

// SectionType represents the section node type.
const SectionType = "Section"

// Section holds data about a specific portion of the report.
// The Level property represents the nesting of the item.
type Section struct {
	Title    string
	Level    int
	id       string
	children []Node
}

// NewSection instantiates a new section with the given data.
func NewSection(id, title string, level int) *Section {
	return &Section{
		id:       id,
		Title:    title,
		Level:    level,
		children: make([]Node, 0),
	}
}

func (s *Section) ID() string           { return s.id }
func (*Section) Type() string           { return SectionType }
func (s *Section) Append(nodes ...Node) { s.children = append(s.children, nodes...) }
func (s *Section) Children() []Node     { return s.children }
