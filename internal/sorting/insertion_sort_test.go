package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 3, 8, 1, 2}, []int{1, 2, 3, 5, 8}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 3, 2, 1, 2}, []int{1, 2, 2, 3, 3}},
	}
	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)
		insertionSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("insertionSort(%v) = %v; want %v", test.input, inputCopy, test.expected)
		}
	}
}

func TestReverseInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{2, 1}},
		{[]int{5, 3, 8, 1, 2}, []int{8, 5, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{3, 3, 2, 1, 2}, []int{3, 3, 2, 2, 1}},
	}
	for _, test := range tests {
		inputCopy := make([]int, len(test.input))
		copy(inputCopy, test.input)
		reverseInsertionSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, test.expected) {
			t.Errorf("reverseInsertionSort(%v) = %v; want %v", test.input, inputCopy, test.expected)
		}
	}
}
