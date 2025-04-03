package data_structures

import "testing"

func StackTestIsEmpty(t *testing.T) {
	stack := &Stack[int]{}
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

func StackTestPushAndSize(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	if stack.Size() != 3 {
		t.Errorf("Expected stack size 3, got %d", stack.Size())
	}
}

func StackTestPeek(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	top, err := stack.Peek()
	if err != nil || top != 30 {
		t.Errorf("Expected top element 30, got %d", top)
	}
}

func StackTestPop(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	popped, err := stack.Pop()
	if err != nil || popped != 30 {
		t.Errorf("Expected popped element 30, got %d", popped)
	}

	if stack.Size() != 2 {
		t.Errorf("Expected stack size 2, got %d", stack.Size())
	}
}

func StackTestPopUntilEmpty(t *testing.T) {
	stack := &Stack[int]{}
	stack.Push(10)
	stack.Push(20)

	_, _ = stack.Pop()
	_, _ = stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

func StackTestPopOnEmptyStack(t *testing.T) {
	stack := &Stack[int]{}
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping from empty stack")
	}
}
