package dynamic

import "testing"

// TestCanSum is a comprehensive test function that covers all test cases for the canSum function
func TestCanSum(t *testing.T) {
	tests := []struct {
		name      string
		targetSum int
		numbers   []int
		expected  bool
	}{
		// Base cases - target sum is zero
		{name: "Target sum is zero with empty array", targetSum: 0, numbers: []int{}, expected: true},
		{name: "Target sum is zero with values", targetSum: 0, numbers: []int{1, 2, 3}, expected: true},

		// Base cases - empty array with non-zero target
		{name: "Target sum is positive with empty array", targetSum: 5, numbers: []int{}, expected: false},

		// Single element array
		{name: "Single element equal to target", targetSum: 5, numbers: []int{5}, expected: true},
		{name: "Single element not equal to target", targetSum: 5, numbers: []int{3}, expected: false},

		// Simple cases with multiple elements
		{name: "Target achievable with two elements", targetSum: 7, numbers: []int{5, 2}, expected: true},
		{name: "Target achievable with multiple elements", targetSum: 8, numbers: []int{2, 3, 5}, expected: true},
		{name: "Target not achievable with multiple elements", targetSum: 7, numbers: []int{2, 4}, expected: false},

		// Cases using repeated elements
		{name: "Target achievable by using one element multiple times", targetSum: 10, numbers: []int{5}, expected: true},
		{name: "Target achievable by using multiple elements multiple times", targetSum: 15, numbers: []int{3, 5}, expected: true},
		{name: "Target not achievable even using repeated elements", targetSum: 11, numbers: []int{5}, expected: false},

		// Cases with different combinations
		{name: "Target 4 with array [3, 4]", targetSum: 4, numbers: []int{3, 4}, expected: true},
		{name: "Target 3 with array [2, 4]", targetSum: 3, numbers: []int{2, 4}, expected: false},
		{name: "Target 0 with array [5, 10]", targetSum: 0, numbers: []int{5, 10}, expected: true},

		// Cases with larger target sum
		{name: "Target 13 with array [6, 9, 3, 8]", targetSum: 13, numbers: []int{6, 9, 3, 8}, expected: false},
		{name: "Target 130 with array [7, 14]", targetSum: 130, numbers: []int{7, 14}, expected: false},
		{name: "Target 12 with array [2, 3]", targetSum: 12, numbers: []int{2, 3}, expected: true},

		// Cases where target is not achievable
		{name: "Target 11 with array [2, 4, 6]", targetSum: 11, numbers: []int{2, 4, 6}, expected: false},
		{name: "Target 300 with array [7, 14]", targetSum: 300, numbers: []int{7, 14}, expected: false},

		// Cases with array containing 1
		{name: "Target 10 with array containing 1", targetSum: 10, numbers: []int{1, 5}, expected: true},
		{name: "Target 7 with array [1]", targetSum: 7, numbers: []int{1}, expected: true},

		// Cases with duplicate values in array
		{name: "Target 6 with array containing duplicates", targetSum: 6, numbers: []int{3, 2, 3}, expected: true},
		{name: "Target 10 with array [5, 5]", targetSum: 10, numbers: []int{5, 5}, expected: true},

		// Edge cases with larger values
		{name: "Target 1 with array [2, 3, 5]", targetSum: 1, numbers: []int{2, 3, 5}, expected: false},
		{name: "Target 100 with array [25]", targetSum: 100, numbers: []int{25}, expected: true},
		{name: "Target 50 with array [7, 8, 9]", targetSum: 50, numbers: []int{7, 8, 9}, expected: true},

		// Cases requiring multiple uses of the same number
		{name: "Target 20 with array [2]", targetSum: 20, numbers: []int{2}, expected: true},
		{name: "Target 15 with array [4, 5]", targetSum: 15, numbers: []int{4, 5}, expected: true},
		{name: "Target 8 with array [2, 3]", targetSum: 8, numbers: []int{2, 3}, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := canSum(test.targetSum, test.numbers)
			if result != test.expected {
				t.Errorf("canSum(%d, %v) = %v; expected %v",
					test.targetSum, test.numbers, result, test.expected)
			}
		})
	}
}

// TestMemoCanSum is a comprehensive test function for the memoized version of canSum
func TestMemoCanSum(t *testing.T) {
	tests := []struct {
		name      string
		targetSum int
		numbers   []int
		expected  bool
	}{
		// Base cases - target sum is zero
		{name: "Target sum is zero with empty array", targetSum: 0, numbers: []int{}, expected: true},
		{name: "Target sum is zero with values", targetSum: 0, numbers: []int{1, 2, 3}, expected: true},

		// Base cases - empty array with non-zero target
		{name: "Target sum is positive with empty array", targetSum: 5, numbers: []int{}, expected: false},

		// Single element array
		{name: "Single element equal to target", targetSum: 5, numbers: []int{5}, expected: true},
		{name: "Single element not equal to target", targetSum: 5, numbers: []int{3}, expected: false},

		// Simple cases with multiple elements
		{name: "Target achievable with two elements", targetSum: 7, numbers: []int{5, 2}, expected: true},
		{name: "Target achievable with multiple elements", targetSum: 8, numbers: []int{2, 3, 5}, expected: true},
		{name: "Target not achievable with multiple elements", targetSum: 7, numbers: []int{2, 4}, expected: false},

		// Cases using repeated elements
		{name: "Target achievable by using one element multiple times", targetSum: 10, numbers: []int{5}, expected: true},
		{name: "Target achievable by using multiple elements multiple times", targetSum: 15, numbers: []int{3, 5}, expected: true},
		{name: "Target not achievable even using repeated elements", targetSum: 11, numbers: []int{5}, expected: false},

		// Cases with different combinations
		{name: "Target 4 with array [3, 4]", targetSum: 4, numbers: []int{3, 4}, expected: true},
		{name: "Target 3 with array [2, 4]", targetSum: 3, numbers: []int{2, 4}, expected: false},
		{name: "Target 0 with array [5, 10]", targetSum: 0, numbers: []int{5, 10}, expected: true},

		// Cases with larger target sum
		{name: "Target 13 with array [6, 9, 3, 8]", targetSum: 13, numbers: []int{6, 9, 3, 8}, expected: false},
		{name: "Target 130 with array [7, 14]", targetSum: 130, numbers: []int{7, 14}, expected: false},
		{name: "Target 12 with array [2, 3]", targetSum: 12, numbers: []int{2, 3}, expected: true},

		// Cases where target is not achievable
		{name: "Target 11 with array [2, 4, 6]", targetSum: 11, numbers: []int{2, 4, 6}, expected: false},
		{name: "Target 300 with array [7, 14]", targetSum: 300, numbers: []int{7, 14}, expected: false},

		// Cases with array containing 1
		{name: "Target 10 with array containing 1", targetSum: 10, numbers: []int{1, 5}, expected: true},
		{name: "Target 7 with array [1]", targetSum: 7, numbers: []int{1}, expected: true},

		// Cases with duplicate values in array
		{name: "Target 6 with array containing duplicates", targetSum: 6, numbers: []int{3, 2, 3}, expected: true},
		{name: "Target 10 with array [5, 5]", targetSum: 10, numbers: []int{5, 5}, expected: true},

		// Edge cases with larger values
		{name: "Target 1 with array [2, 3, 5]", targetSum: 1, numbers: []int{2, 3, 5}, expected: false},
		{name: "Target 100 with array [25]", targetSum: 100, numbers: []int{25}, expected: true},
		{name: "Target 50 with array [7, 8, 9]", targetSum: 50, numbers: []int{7, 8, 9}, expected: true},

		// Cases requiring multiple uses of the same number
		{name: "Target 20 with array [2]", targetSum: 20, numbers: []int{2}, expected: true},
		{name: "Target 15 with array [4, 5]", targetSum: 15, numbers: []int{4, 5}, expected: true},
		{name: "Target 8 with array [2, 3]", targetSum: 8, numbers: []int{2, 3}, expected: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := memoCanSum(test.targetSum, test.numbers, make(map[int]struct{}))
			if result != test.expected {
				t.Errorf("memoCanSum(%d, %v) = %v; expected %v",
					test.targetSum, test.numbers, result, test.expected)
			}
		})
	}
}
