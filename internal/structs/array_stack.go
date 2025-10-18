package structs

type ArrayStack[T any] struct {
	data []T
	size int // Tracks the current size
}

// Push adds an element to the top of the stack
// Time Complexity: O(1) amortized
func (as *ArrayStack[T]) Push(value T) {
	// If there's no capacity, grow the slice
	if as.size >= len(as.data) {
		// Growth strategy: double the capacity (minimum 8)
		newCap := len(as.data) * 2
		if newCap == 0 {
			newCap = 8
		}
		newData := make([]T, newCap)
		copy(newData, as.data)
		as.data = newData
	}
	as.data[as.size] = value
	as.size++
}

// Pop removes and returns the top element from the stack
// Time Complexity: O(1)
// Returns the element and a boolean indicating if the stack was not empty
func (as *ArrayStack[T]) Pop() (T, bool) {
	if as.size == 0 {
		var zero T
		return zero, false
	}
	as.size--
	value := as.data[as.size]
	// Clear the reference to allow GC
	var zero T
	as.data[as.size] = zero
	return value, true
}

// Peek returns the top element without removing it
func (as *ArrayStack[T]) Peek() (T, bool) {
	if as.size == 0 {
		var zero T
		return zero, false
	}
	return as.data[as.size-1], true
}

// Size returns the number of elements in the stack
func (as *ArrayStack[T]) Size() int {
	return as.size
}

// IsEmpty indicates if the stack is empty
func (as *ArrayStack[T]) IsEmpty() bool {
	return as.size == 0
}
