package main

import (
	"fmt"

	"github.com/betofloresbaca/algorithms/internal/other"
	"github.com/betofloresbaca/algorithms/internal/sorting"
)

func main() {
	fmt.Println("=== Algorithms Demo ===")
	fmt.Println()

	fmt.Println("1. Insertion Sort:")
	sorting.InsertionSortMain()

	fmt.Println("\n2. Merge Sort:")
	sorting.MergeSortMain()

	fmt.Println("\n3. Linear Search:")
	other.LinearSearchMain()

	fmt.Println("\n4. Bit Array Sum:")
	other.BitArraySumMain()
}
