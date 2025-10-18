package dynamic

import "testing"

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=3", 3, 2},
		{"n=4", 4, 3},
		{"n=5", 5, 5},
		{"n=6", 6, 8},
		{"n=7", 7, 13},
		{"n=8", 8, 21},
		{"n=10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("fibonacci(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(10)
	}
}

func TestMemoryFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"n=0", 0, 0},
		{"n=1", 1, 1},
		{"n=2", 2, 1},
		{"n=3", 3, 2},
		{"n=4", 4, 3},
		{"n=5", 5, 5},
		{"n=6", 6, 8},
		{"n=7", 7, 13},
		{"n=8", 8, 21},
		{"n=10", 10, 55},
		{"n=15", 15, 610},
		{"n=20", 20, 6765},
		{"n=100", 100, 3736710778780434371},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := memoFibonacci(tt.input, make(map[int]int))
			if result != tt.expected {
				t.Errorf("MemoryFibonacci(%d) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkMemoryFibonacci(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		memoFibonacci(20, make(map[int]int))
	}
}
