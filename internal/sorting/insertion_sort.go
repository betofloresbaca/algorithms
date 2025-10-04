package sorting

import "fmt"

func insertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func reverseInsertionSort(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for j >= 0 && data[j] < key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

func InsertionSortMain() {
	data := []int{13, 3, 12, 8, 37, 4, 9, 2, 15, 20}
	insertionSort(data)
	// reverseInsertionSort(data)
	fmt.Printf("Insertion Sort: %v\n", data)
}
