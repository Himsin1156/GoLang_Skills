package main

import "fmt"

// Operation using function

// func addnum(num1 int, num2 int) int {
// 	return num1 + num2
// }

// func minusnum(num1 int, num2 int) int {
// 	return num1 - num2

// }

// func multiplicationNum(num1 int, num2 int) int {
// 	return num1 * num2
// }

// func divisionNum(num1 int, num2 int) int {
// 	return num1 / num2
// }

// 2. Multiple / Many return values:

func hobbies() (string, string, string, string, string, int) {

	return "Listenning Music", "Playing Guitar", "Hanging out on trip", "Name = Himanshu", "Age =", 30
}

func main() {

	// resultAdd := addnum(10, 2)
	// fmt.Println("Sum for given numbers are ", resultAdd)

	// resultMinus := minusnum(10, 2)
	// fmt.Println("Substraction of given numbers are", resultMinus)

	// resultMultiplication := multiplicationNum(10, 2)
	// fmt.Println("Multiplication of given numbers are", resultMultiplication)

	// resultDivision := divisionNum(10, 2)
	// fmt.Println("Division of given numbers are", resultDivision)

	// 2. multiple returning values:

	//1. fmt.Println(hobbies())
	// 2. multiple way to assign function to any variable and use that to print / output
	hobby1, hobb2, hobby3, name, age, number := hobbies()

	fmt.Println(hobby1, hobb2, hobby3, name, age, number)

}
