package BinarySearch

import "fmt"

// create Root node and insert child nodes
func BinarySearchTreeOperation() {
	fmt.Println("Binary Search Tree Assignment")
	test := NewBst()
	// Inserting some data
	test.Insert(10)
	test.Insert(20)
	test.Insert(40)
	test.Insert(60)
	test.Insert(30)
	test.Insert(50)

	// InOrder operation
	InOrder(test.Root)
	fmt.Println()

	// // PreOrder operation
	PreOrder(test.Root)
	fmt.Println()

	// PostOrder operation
	PostOrder(test.Root)
	fmt.Println()
}
