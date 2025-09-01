package main

import "fmt"

// 2. function using range to use slice to sum the numbers
func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total = total + num
	}

	return total
}

func main() {

	//1. fmt.Println(1, 2, 3, 4, 5, "Himan", 30)
	//2. Sum of N number or range of multiple numbers
	// slice
	nums := []int{100, 200, 200, 500} //3. initializzing a slice in nums

	//result := sum(100, 2, 3, 4, 5)
	result := sum(nums...) // passing argument for slice nums with ... as multiple numbers
	fmt.Println(result)

}
