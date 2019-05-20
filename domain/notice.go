package domain

// Notice represents a notice text not tied to a speaker.
type Notice struct {
	Content string
}

// NewNotice instantiates a new notice with given content.
func NewNotice(content string) *Notice {
	return &Notice{
		Content: content,
	}
}

func (*Notice) Append(nodes ...Node) {}
func (*Notice) Children() []Node     { return []Node{} }
