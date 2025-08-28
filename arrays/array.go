package main

import "fmt"

// Arrays

func main() {

	// arr1 := [4]int{11, 22, 33, 44}   // single dimensional array

	arr2 := [3][3]int{{12, 13, 14}, {22, 23, 24}, {32, 33, 34}} // two dimensional array   .. 3 aarats with 3-3-3 elements

	// way to retrieve the output from 2d array
	fmt.Println(arr2)
	fmt.Println(arr2[0][2])
	fmt.Println(arr2[1][0])
	fmt.Println(arr2[2][1])

	// We can use array if below point are clear and needed

	// - Fixed size - that is pridictable
	// - Memory Optimization
	// - Contact time access

	// - We use slices majorly in golang

}
