package data_structures

import (
	"errors"
)

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

// Enqueue adds an element to the end of the queue
// O(1) time
// O(1) memory
func (q *Queue[T]) Enqueue(item T) {
	// We create the new node
	new_node := &Node[T]{}
	new_node.value = item

	// If this is the first time we add a new node
	if q.head == nil && q.tail == nil {
		q.head = new_node
		q.tail = new_node
		// If this is NOT the first time we add a new node
	} else {
		// We just add it to the end
		q.tail.next = new_node
		// We update our tail
		q.tail = new_node
	}
	q.size++
}

// Dequeue removes and returns the first element of the queue
// O(1) time
// O(1) memory
func (q *Queue[T]) Dequeue() (T, error) {
	// If the queue is emprt, we can't dequeue
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("ERROR: Can't dequeue because stack is empty")
	}

	// We take out the node at the front
	dequeued_node := q.head
	// We update our head
	q.head = q.head.next

	// If our head ever becomes empty, then we should also
	// make the tail empty as well
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return dequeued_node.value, nil
}

// Peek returns the first element without removing it
// O(1) time
// O(1) memory
func (q *Queue[T]) Peek() (any, error) {
	// If the queue is emprt, we can't peek
	if q.IsEmpty() {
		return nil, errors.New("ERROR: Can't peek because stack is empty")
	}
	return q.head.value, nil
}

// IsEmpty checks if the queue is empty
// O(1) time
// O(1) memory
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}
