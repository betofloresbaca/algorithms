package sorting

func mergeSort(data []int) []int {
	// Base case
	if len(data) <= 1 {
		return data
	}
	// Resursive case
	// Sort left and right part
	middle := len(data) / 2
	leftPart := mergeSort(data[:middle])
	rightPart := mergeSort(data[middle:])
	// Merge the results in the results array
	result := make([]int, len(data))
	leIdx, riIdx, reIdx := 0, 0, 0
	for reIdx < len(data) {
		if leIdx >= len(leftPart) {
			result[reIdx] = rightPart[riIdx]
			riIdx++
		} else if riIdx >= len(rightPart) {
			result[reIdx] = leftPart[leIdx]
			leIdx++
		} else if leftPart[leIdx] > rightPart[riIdx] {
			result[reIdx] = rightPart[riIdx]
			riIdx++
		} else {
			result[reIdx] = leftPart[leIdx]
			leIdx++
		}
		reIdx++
	}
	return result
}
