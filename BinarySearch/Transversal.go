package BinarySearch

import "fmt"

// Print the tree in-order
// Traverse the left sub-tree, root, right sub-tree
func InOrder(root *Node) {
	if root != nil {
		InOrder(root.left)
		fmt.Printf("%d ", root.value)
		InOrder(root.right)
	}
}

// Print the tree pre-order
// Traverse the root, left sub-tree, right sub-tree
func PreOrder(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.value)
		PreOrder(root.left)
		PreOrder(root.right)
	}
}

// Print the tree post-order
// Traverse left sub-tree, right sub-tree, root
func PostOrder(root *Node) {
	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Printf("%d ", root.value)
	}
}
