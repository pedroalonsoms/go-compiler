package data_structures

import (
	"errors"
)

// This is a Stack data structure implementation
type Stack[T any] struct {
	top  *Node[T]
	size int
}

// Push adds an element to the top of the stack
// O(1) time
// O(1) memory
func (s *Stack[T]) Push(item T) {
	// We create a new node that points to our top node
	// ... <- o <- o <- top <- new_node
	new_node := &Node[T]{}
	new_node.value = item
	new_node.next = s.top

	// Now our top node is the newly-added node
	s.top = new_node
	s.size++
}

// Pop removes and returns the top element of the stack
// O(1) time
// O(1) memory
func (s *Stack[T]) Pop() (T, error) {
	// If it is empty then we can't pop
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("ERROR: Can't pop because stack is empty")
	}

	// We just make it so that our top node points to the next node
	popped_node := s.top
	s.top = s.top.next
	s.size--

	return popped_node.value, nil
}

// Peek returns the top element without removing it
// O(1) time
// O(1) memory
func (s *Stack[T]) Peek() (T, error) {
	// If it is empty then we can't peek
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("ERROR: Can't peek because stack is empty")
	}

	return s.top.value, nil
}

// IsEmpty checks if the stack is empty
// O(1) time
// O(1) memory
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// Size returns the number of elements in the stack
// O(1) time
// O(1) memory
func (s *Stack[T]) Size() int {
	return s.size
}
