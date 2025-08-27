package main

import "fmt"

func main() {

	//Main syntax for declaring the variable:

	//var Name string = "Himanshu"
	//var Phone number ="0761593747"
	// var Name = "Himanshu"
	// var phone int = 761593747
	// var isadult = true
	// var height = 6.1
	// var age int = 30

	//Now below we will use shorthand syntax:

	Name := "Himanshu"
	phone := 761593747
	isadult := true
	height := 6.1
	age := 30
	// Way to define const ----- Const cannot be refined or change
	const post string = "Senior Software Engineer"
	// Way to define multiple Const variables
	const (
		surname  = "Singh"
		Office   = "Netinsight AB"
		Location = "Stockholm Sweden"
	)

	fmt.Println(
		"Name of candidate ", Name, surname,
		" ,", "Mobile No. ", phone,
		", ", "Older than 18 ", isadult,
		" ,", "Height is ", height,
		" ,", "Age is ", age,
		" ,", " Profile is ", post,
		" ,", " Office going to ", Office,
		" ", "Location is ", Location)

}
