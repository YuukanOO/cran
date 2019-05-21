package domain

// Intervention represents a speaker allocution.
type Intervention struct {
	id        string
	SpeakerID string
	Content   string
}

// NewIntervention instantiates a new intervention with given data.
func NewIntervention(ID, speakerID, content string) *Intervention {
	return &Intervention{
		id:        ID,
		SpeakerID: speakerID,
		Content:   content,
	}
}

func (*Intervention) Append(nodes ...Node) {}
func (*Intervention) Children() []Node     { return []Node{} }
func (*Intervention) Type() string         { return "Intervention" }
func (i *Intervention) ID() string         { return i.id }
