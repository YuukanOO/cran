package cran

type Root struct {
	children []Node
}

func newRoot() *Root {
	return &Root{
		children: make([]Node, 0),
	}
}

func (r *Root) Append(n Node) { r.children = append(r.children, n) }
func (r *Root) Children() []Node { return r.children }

type Node interface {
	Append(n Node)
	Children() []Node
}

type SectionNode struct{
	Title string
	children []Node
}

func (r *SectionNode) Append(n Node) { r.children = append(r.children, n) }
func (r *SectionNode) Children() []Node { return r.children }

func newSection(title string) *SectionNode {
	return &SectionNode{
		Title: title,
		children: make([]Node, 0),
	}
}

type InterventionNode struct {
	SpeakerID string
	Sentence string
}

func (r *InterventionNode) Append(n Node) { }
func (r *InterventionNode) Children() []Node { return nil }

func newIntervention(speaker, sentence string) *InterventionNode {
	return &InterventionNode{
		SpeakerID: speaker,
		Sentence: sentence,
	}
}

type assembleNationaleCRProvider struct{}

func (p *assembleNationaleCRProvider) Fetch(URL string, callback ProviderCallback) {
	callback(nil, nil)
}