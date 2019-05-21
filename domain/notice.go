package domain

// Notice represents a notice text not tied to a speaker.
type Notice struct {
	id      string
	Content string
}

// NewNotice instantiates a new notice with given content.
func NewNotice(ID, content string) *Notice {
	return &Notice{
		id:      ID,
		Content: content,
	}
}

func (*Notice) Append(nodes ...Node) {}
func (*Notice) Children() []Node     { return []Node{} }
func (*Notice) Type() string         { return "Notice" }
func (n *Notice) ID() string         { return n.id }
