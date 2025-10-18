package structs

import (
	"testing"
)

// TestArrayStack_Push tests adding elements to the stack
func TestArrayStack_Push(t *testing.T) {
	stack := &ArrayStack[int]{}

	stack.Push(1)
	if stack.Size() != 1 {
		t.Errorf("expected size 1, got %d", stack.Size())
	}

	stack.Push(2)
	stack.Push(3)
	if stack.Size() != 3 {
		t.Errorf("expected size 3, got %d", stack.Size())
	}
}

// TestArrayStack_Pop tests removing elements from the stack
func TestArrayStack_Pop(t *testing.T) {
	stack := &ArrayStack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Pop from the top (LIFO)
	val, ok := stack.Pop()
	if !ok || val != 3 {
		t.Errorf("expected pop(3, true), got (%d, %v)", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != 2 {
		t.Errorf("expected pop(2, true), got (%d, %v)", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != 1 {
		t.Errorf("expected pop(1, true), got (%d, %v)", val, ok)
	}

	if stack.Size() != 0 {
		t.Errorf("expected size 0, got %d", stack.Size())
	}
}

// TestArrayStack_PopEmpty tests Pop on an empty stack
func TestArrayStack_PopEmpty(t *testing.T) {
	stack := &ArrayStack[int]{}

	val, ok := stack.Pop()
	if ok {
		t.Errorf("expected pop(_, false), got (_, true)")
	}

	if val != 0 { // zero value for int
		t.Errorf("expected zero value, got %d", val)
	}
}

// TestArrayStack_Peek tests reading without removal
func TestArrayStack_Peek(t *testing.T) {
	stack := &ArrayStack[int]{}
	stack.Push(10)
	stack.Push(20)

	val, ok := stack.Peek()
	if !ok || val != 20 {
		t.Errorf("expected peek(20, true), got (%d, %v)", val, ok)
	}

	// Size should be the same after Peek
	if stack.Size() != 2 {
		t.Errorf("expected size 2 after peek, got %d", stack.Size())
	}
}

// TestArrayStack_PeekEmpty tests Peek on an empty stack
func TestArrayStack_PeekEmpty(t *testing.T) {
	stack := &ArrayStack[int]{}

	val, ok := stack.Peek()
	if ok {
		t.Errorf("expected peek(_, false), got (_, true)")
	}

	if val != 0 {
		t.Errorf("expected zero value, got %d", val)
	}
}

// TestArrayStack_IsEmpty tests empty state verification
func TestArrayStack_IsEmpty(t *testing.T) {
	stack := &ArrayStack[int]{}

	if !stack.IsEmpty() {
		t.Errorf("expected IsEmpty()=true, got false")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("expected IsEmpty()=false, got true")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Errorf("expected IsEmpty()=true after pop, got false")
	}
}

// TestArrayStack_Size tests the size counter
func TestArrayStack_Size(t *testing.T) {
	stack := &ArrayStack[int]{}

	if stack.Size() != 0 {
		t.Errorf("expected size 0, got %d", stack.Size())
	}

	for i := 1; i <= 5; i++ {
		stack.Push(i)
		if stack.Size() != i {
			t.Errorf("expected size %d, got %d", i, stack.Size())
		}
	}
}

// TestArrayStack_GrowthStrategy tests that the stack grows correctly
func TestArrayStack_GrowthStrategy(t *testing.T) {
	stack := &ArrayStack[int]{}

	// Push more elements than initial capacity (8)
	for i := 1; i <= 20; i++ {
		stack.Push(i)
	}

	if stack.Size() != 20 {
		t.Errorf("expected size 20, got %d", stack.Size())
	}

	// Verify LIFO
	val, _ := stack.Pop()
	if val != 20 {
		t.Errorf("expected last element 20, got %d", val)
	}
}

// TestArrayStack_WithStrings tests with string generic type
func TestArrayStack_WithStrings(t *testing.T) {
	stack := &ArrayStack[string]{}

	stack.Push("hello")
	stack.Push("world")

	val, ok := stack.Pop()
	if !ok || val != "world" {
		t.Errorf("expected pop('world', true), got ('%s', %v)", val, ok)
	}

	val, ok = stack.Pop()
	if !ok || val != "hello" {
		t.Errorf("expected pop('hello', true), got ('%s', %v)", val, ok)
	}
}

// TestArrayStack_LIFO verifies complete LIFO behavior
func TestArrayStack_LIFO(t *testing.T) {
	stack := &ArrayStack[int]{}
	values := []int{5, 3, 8, 1, 9}

	for _, v := range values {
		stack.Push(v)
	}

	// We should get values in reverse order
	for i := len(values) - 1; i >= 0; i-- {
		val, ok := stack.Pop()
		if !ok || val != values[i] {
			t.Errorf("expected pop(%d, true), got (%d, %v)", values[i], val, ok)
		}
	}

	if !stack.IsEmpty() {
		t.Errorf("expected stack to be empty after all pops")
	}
}
