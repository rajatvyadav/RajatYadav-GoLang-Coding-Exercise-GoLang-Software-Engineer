package BinarySearch

type Node struct {
	left  *Node
	right *Node
	value int
	// Add whatever you want here
	//...
	//...
}

type Bst struct {
	Root *Node
}

// Construtor
func NewBst() *Bst {
	bst := new(Bst)
	return bst
}
