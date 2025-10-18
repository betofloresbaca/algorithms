package dynamic

func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return make([]int, 0)
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if innerResult := howSum(targetSum-num, numbers); innerResult != nil {
			return append([]int{num}, innerResult...)
		}
	}
	return nil
}

func memoHowSum(targetSum int, numbers []int, cache map[int]struct{}) []int {
	if _, exists := cache[targetSum]; exists {
		return nil
	}
	if targetSum == 0 {
		return make([]int, 0)
	}
	if targetSum < 0 {
		return nil
	}
	for _, num := range numbers {
		if innerResult := memoHowSum(targetSum-num, numbers, cache); innerResult != nil {
			return append(innerResult, num)
		}
	}
	cache[targetSum] = struct{}{}
	return nil
}
