package other

import "fmt"

func bitArraySum(a []int, b []int) []int {
	result := make([]int, len(a)+1)
	carry := 0
	for i := 0; i < len(a); i++ {
		sum := a[i] + b[i] + carry
		result[i] = sum % 2
		carry = sum / 2
	}
	result[len(a)] = carry
	return result
}

func BitArraySumMain() {
	a := []int{1, 0, 1, 1, 1, 0, 0, 0}
	b := []int{1, 1, 0, 0, 0, 1, 1, 1}
	c := bitArraySum(a, b)
	fmt.Printf("Bit Array Sum:\n")
	fmt.Printf("A: %v\n", a)
	fmt.Printf("B: %v\n", b)
	fmt.Printf("Sum: %v\n", c)
}
