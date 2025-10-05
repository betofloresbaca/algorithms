package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
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
		result := mergeSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("mergeSort(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
