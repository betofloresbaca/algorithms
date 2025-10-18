package dynamic

func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, value := range numbers {
		if canSum(targetSum-value, numbers) {
			return true
		}
	}
	return false
}

func memoCanSum(targetSum int, numbers []int, memo map[int]struct{}) bool {
	if _, exists := memo[targetSum]; exists {
		return false
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, value := range numbers {
		if memoCanSum(targetSum-value, numbers, memo) {
			return true
		}

	}
	memo[targetSum] = struct{}{}
	return false

}
