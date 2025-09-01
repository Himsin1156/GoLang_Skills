package main

import (
	"fmt"
)

// Range is used for iteration in data structure.

func main() {

	// 1. Range on slices

	//num1 := []int{11, 22, 33, 44, 55}

	// for i := 0; i < len(num1); i++ {
	//  	fmt.Println(num1[i])

	// }
	//sum := 0
	// for v, num1 := range num1 {
	// 	fmt.Println(num1, v)
	// 	//sum = sum + num1
	// 	//fmt.Println(num1)
	// }
	//fmt.Println(sum)

	// 2. Range interation on Maps

	// m := map[string]string{"Name": "Himashu", "Work": "Ingenjor"}

	// for _, val := range m { // k = key , v =value
	// 	fmt.Println(val)
	// }

	// 3.Range on string   - Unicode code point rune (rune is data structure)

	for i, unicode := range "Himanshu" {
		fmt.Println(unicode, " ", i, " ", string(unicode))
	}

}
