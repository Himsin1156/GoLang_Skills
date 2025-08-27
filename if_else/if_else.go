package main

import "fmt"

var age int

func main() {

	fmt.Println("Enter you age")
	fmt.Scanln(&age)

	if age >= 18 {
		fmt.Println("Eligible for vote as the person is adult")
	} else if age >= 12 {
		fmt.Println("Not Eligible for voting as less the criteria of age i.e 18 years")
	} else {
		fmt.Println("Go home and see pogo and wait for your time...Hahaha")
	}

}
