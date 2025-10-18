package dynamic

import "testing"

// TestGridTraveler is a comprehensive test function that covers all test cases
func TestGridTraveler(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		m        int
		expected int
	}{
		// Base cases - zero dimensions
		{name: "Grid with zero width", n: 0, m: 5, expected: 0},
		{name: "Grid with zero height", n: 5, m: 0, expected: 0},
		{name: "Grid both dimensions zero", n: 0, m: 0, expected: 0},

		// Base cases - grid 1xN or Nx1
		{name: "Grid 1x1", n: 1, m: 1, expected: 1},
		{name: "Grid 1x2", n: 1, m: 2, expected: 1},
		{name: "Grid 2x1", n: 2, m: 1, expected: 1},
		{name: "Grid 1x10", n: 1, m: 10, expected: 1},
		{name: "Grid 10x1", n: 10, m: 1, expected: 1},

		// Square grids
		{name: "Grid 2x2", n: 2, m: 2, expected: 2},
		{name: "Grid 3x3", n: 3, m: 3, expected: 6},
		{name: "Grid 4x4", n: 4, m: 4, expected: 20},
		{name: "Grid 5x5", n: 5, m: 5, expected: 70},

		// Rectangular grids
		{name: "Grid 2x3", n: 2, m: 3, expected: 3},
		{name: "Grid 3x2", n: 3, m: 2, expected: 3},
		{name: "Grid 2x4", n: 2, m: 4, expected: 4},
		{name: "Grid 3x4", n: 3, m: 4, expected: 10},
		{name: "Grid 4x3", n: 4, m: 3, expected: 10},

		// Commutativity tests
		{name: "Commutative: 3 and 5", n: 3, m: 5, expected: 15},
		{name: "Commutative: 5 and 3", n: 5, m: 3, expected: 15},
		{name: "Commutative: 4 and 2", n: 4, m: 2, expected: 4},
		{name: "Commutative: 2 and 4", n: 2, m: 4, expected: 4},

		// Large values
		// {name: "Grid 18x18", n: 18, m: 18, expected: 2333606220},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := gridTraveler(test.n, test.m)
			if result != test.expected {
				t.Errorf("gridTraveler(%d, %d) = %d; expected %d",
					test.n, test.m, result, test.expected)
			}
			// Verify non-negativity
			if result < 0 {
				t.Errorf("gridTraveler(%d, %d) returned %d; should be non-negative",
					test.n, test.m, result)
			}
		})
	}
}

// TestMemoizedGridTraveler is a comprehensive test function for the memoized version
func TestMemoizedGridTraveler(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		m        int
		expected int
	}{
		// Base cases - zero dimensions
		{name: "Grid with zero width", n: 0, m: 5, expected: 0},
		{name: "Grid with zero height", n: 5, m: 0, expected: 0},
		{name: "Grid both dimensions zero", n: 0, m: 0, expected: 0},

		// Base cases - grid 1xN or Nx1
		{name: "Grid 1x1", n: 1, m: 1, expected: 1},
		{name: "Grid 1x2", n: 1, m: 2, expected: 1},
		{name: "Grid 2x1", n: 2, m: 1, expected: 1},
		{name: "Grid 1x10", n: 1, m: 10, expected: 1},
		{name: "Grid 10x1", n: 10, m: 1, expected: 1},

		// Square grids
		{name: "Grid 2x2", n: 2, m: 2, expected: 2},
		{name: "Grid 3x3", n: 3, m: 3, expected: 6},
		{name: "Grid 4x4", n: 4, m: 4, expected: 20},
		{name: "Grid 5x5", n: 5, m: 5, expected: 70},

		// Rectangular grids
		{name: "Grid 2x3", n: 2, m: 3, expected: 3},
		{name: "Grid 3x2", n: 3, m: 2, expected: 3},
		{name: "Grid 2x4", n: 2, m: 4, expected: 4},
		{name: "Grid 3x4", n: 3, m: 4, expected: 10},
		{name: "Grid 4x3", n: 4, m: 3, expected: 10},

		// Commutativity tests
		{name: "Commutative: 3 and 5", n: 3, m: 5, expected: 15},
		{name: "Commutative: 5 and 3", n: 5, m: 3, expected: 15},
		{name: "Commutative: 4 and 2", n: 4, m: 2, expected: 4},
		{name: "Commutative: 2 and 4", n: 2, m: 4, expected: 4},

		// Large values
		{name: "Grid 18x18", n: 18, m: 18, expected: 2333606220},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := memoGridTraveler(test.n, test.m, make(map[dims]int))
			if result != test.expected {
				t.Errorf("memoGridTraveler(%d, %d) = %d; expected %d",
					test.n, test.m, result, test.expected)
			}
			// Verify non-negativity
			if result < 0 {
				t.Errorf("memoGridTraveler(%d, %d) returned %d; should be non-negative",
					test.n, test.m, result)
			}
		})
	}
}
