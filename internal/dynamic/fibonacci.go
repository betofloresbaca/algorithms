package dynamic

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func memoFibonacci(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	result := memoFibonacci(n-1, memo) + memoFibonacci(n-2, memo)
	memo[n] = result
	return result
}
