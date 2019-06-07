package generating

// InterventionType represents the type of an intervention.
const InterventionType = "Intervention"

// Intervention represents a speaker allocution.
type Intervention struct {
	id        string
	SpeakerID string
	Content   string
}

// NewIntervention instantiates a new intervention with given data.
func NewIntervention(id, speakerID, content string) *Intervention {
	return &Intervention{
		id:        id,
		SpeakerID: speakerID,
		Content:   content,
	}
}

func (*Intervention) Append(nodes ...Node) {}
func (*Intervention) Children() []Node     { return []Node{} }
func (*Intervention) Type() string         { return InterventionType }
func (i *Intervention) ID() string         { return i.id }
