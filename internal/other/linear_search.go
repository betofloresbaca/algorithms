package other

import "fmt"

func linearSearch(data []int, value int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == value {
			return i
		}
	}
	return -1
}

func LinearSearchMain() {
	data := []int{13, 3, 12, 8, 37, 4, 9, 2, 15, 20}
	index := linearSearch(data, 2)
	fmt.Printf("Linear Search - Looking for 2 in %v\n", data)
	fmt.Printf("Found at index: %v\n", index)
}
