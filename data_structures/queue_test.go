package data_structures

import "testing"

func QueueTestEnqueue(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(10)
	q.Enqueue(20)
	if q.size != 2 {
		t.Errorf("Expected queue size 2, got %d", q.size)
	}
}

func QueueTestDequeue(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(10)
	q.Enqueue(20)
	item, err := q.Dequeue()
	if err != nil || item != 10 {
		t.Errorf("Expected 10, got %v", item)
	}
	if q.size != 1 {
		t.Errorf("Expected queue size 1, got %d", q.size)
	}
}

func QueueTestDequeueEmpty(t *testing.T) {
	q := &Queue[int]{}
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing from an empty queue")
	}
}

func QueueTestPeek(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(30)
	item, err := q.Peek()
	if err != nil || item != 30 {
		t.Errorf("Expected 30, got %v", item)
	}
}

func QueueTestPeekEmpty(t *testing.T) {
	q := &Queue[int]{}
	_, err := q.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking into an empty queue")
	}
}

func QueueTestIsEmpty(t *testing.T) {
	q := &Queue[int]{}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
	q.Enqueue(40)
	if q.IsEmpty() {
		t.Errorf("Expected queue to be non-empty")
	}
}

func QueueTestSize(t *testing.T) {
	q := &Queue[int]{}
	q.Enqueue(50)
	q.Enqueue(60)
	q.Enqueue(70)
	if q.size != 3 {
		t.Errorf("Expected queue size 3, got %d", q.size)
	}
}
