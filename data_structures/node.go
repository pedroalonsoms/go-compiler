package data_structures

// This is a node we need to emulate a linked list
type Node[T any] struct {
	value T
	next  *Node[T]
}
