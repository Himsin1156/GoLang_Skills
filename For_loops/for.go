package main

import "fmt"

var num int

func main() {

	// fmt.Println("Enter a number to form a table")
	// fmt.Scanln(&num)
	// for i := 1; i <= 10; i++ {

	// 	fmt.Println("Table of ", num, " *", i, " = ", num*i)

	// }

	//While loop using for as while constuct is not available in GoLang
	// i := 1

	// for i <= 100 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// New Feature in GoLang1.22 i.e Range see below

	for i := range 11 {
		fmt.Println(i)
	}

}
