package generating

// NoticeType represents the report node type.
const NoticeType = "Notice"

// Notice represents a notice text not tied to a speaker.
type Notice struct {
	id      string
	Content string
}

// NewNotice instantiates a new notice with given content.
func NewNotice(id, content string) *Notice {
	return &Notice{
		id:      id,
		Content: content,
	}
}

func (*Notice) Append(nodes ...Node) {}
func (*Notice) Children() []Node     { return []Node{} }
func (*Notice) Type() string         { return NoticeType }
func (n *Notice) ID() string         { return n.id }
