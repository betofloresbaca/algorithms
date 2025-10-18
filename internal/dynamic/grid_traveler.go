package dynamic

type dims struct {
	n int
	m int
}

// 0,1 -> 0
// 1,0 -> 0
// 1,1 -> 1

func gridTraveler(n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	return gridTraveler(n-1, m) + gridTraveler(n, m-1)
}

func memoGridTraveler(n, m int, memo map[dims]int) int {
	if value, exists := memo[dims{n, m}]; exists {
		return value
	}
	if value, exists := memo[dims{m, n}]; exists {
		return value
	}
	if n == 0 || m == 0 {
		return 0
	}
	if n == 1 || m == 1 {
		return 1
	}
	result := memoGridTraveler(n-1, m, memo) + memoGridTraveler(n, m-1, memo)
	memo[dims{n, m}] = result
	return result
}
