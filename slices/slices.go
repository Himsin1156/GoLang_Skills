package main

import "fmt"

// slices - no need to give length when using slices
// Automatical expandable and most used construct in Go
// Slices -> Dynamic array
// +Useful method

func main() {

	//**************************//
	// 0. this is uninitialize slice is nil (null).
	//**************************//

	//var num []int
	//fmt.Print(num == nil)
	//fmt.Println(len(num))

	// 1. Another and standard way to initialize slice in go via make function i.e  a inbuilt fuction

	//var num = make([]int, 1, 5)

	//fmt.Println(len(num))
	// fmt.Println(num == nil) // this is not nil as a capacity is here.
	// fmt.Println(cap(num))   // Cap function is to show capacity of slice. Capacity -> is maximum element can fit.

	// below is the way to add element in dynamically in slice.
	//num := []int{}
	// num[0] = 5
	//num = append(num, 1) //, 2, 44, 55, 66, 88, 99)
	// num = append(num, 10)
	// fmt.Println(num)
	// fmt.Println(cap(num))

	//*******************//
	// 2. Copy function in slice , if copy of slice is needed then we can use copy function.
	//*******************//

	// var num = make([]int, 0, 5)

	// num = append(num, 1, 2, 3)

	// var num2 = make([]int, len(num))

	// // Copy function in slice to cpy slice.

	// copy(num2, num)

	// fmt.Println(num, num2)

	//*******************//
	// 3. Slice operator
	//*******************//

	// var slic = []int{2, 4, 6, 8}
	// fmt.Println(slic[0:2]) // --- output [2, 4]
	// fmt.Println(slic[0:3]) // --- output [2 , 4 , 6]
	// fmt.Println(slic[:2])  // --- output [2 4] : also Another way to use slice operator by default started slicing from 0
	// fmt.Println(slic[1:])  // --- output [4 6 8] when limit is not given then starting from 1 element and fetch all other.

	// When using slice operator [0,limit] like above starts from 0 and goes to limit n-1.

	//*******************//
	// 4. Slice Package
	//*******************//

	// var range1 = []int{1, 2, 3, 4}
	// var range2 = []int{1, 2, 3, 5}

	// fmt.Println(slices.Equal(range1, range2))

	//**************************//
	// 5. Two Dimensional Slices
	//***************************//

	var twoDimentionArr = [][]int{{11, 22, 33}, {44, 55, 66}}

	fmt.Println(twoDimentionArr[0][2])
	fmt.Println(twoDimentionArr[1][1])

}
