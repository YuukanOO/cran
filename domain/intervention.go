package domain

// Intervention represents a speaker allocution.
type Intervention struct {
	SpeakerID string
	Content   string
}

// NewIntervention instantiates a new intervention with given data.
func NewIntervention(speakerID, content string) *Intervention {
	return &Intervention{
		SpeakerID: speakerID,
		Content:   content,
	}
}

func (*Intervention) Append(node Node) {}
func (*Intervention) Children() []Node { return []Node{} }
