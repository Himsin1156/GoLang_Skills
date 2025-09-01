package main

import "fmt"

func counter() func() int {
	// no passing argument but return a function func() and func() will return integer
	var count int = 0

	return func() int {
		count += 1 // will return increment count
		return count
	}

}

func main() {

	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}
