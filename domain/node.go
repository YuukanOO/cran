package domain

// Node is the base interface to be implemented to manage a tree like structure.
// Since that's exactly what a report is, it makes sense.
type Node interface {
	Append(node Node)
	Children() []Node
}
