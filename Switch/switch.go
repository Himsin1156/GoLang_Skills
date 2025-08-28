package main

import (
	"fmt"
	"time"
)

var day string

func main() {

	// fmt.Println("Enter a day")
	// fmt.Scanln(&day)

	// switch day {
	// case "Monday":
	// 	fmt.Println("Måndag är veckodag")
	// case "Tuesday":
	// 	fmt.Println("Tisdag är veckodag")
	// case "Wednesday":
	// 	fmt.Println("Onsdag är veckodag")
	// case "Thursday":
	// 	fmt.Println("Torsdag är veckodag")
	// case "Friday":
	// 	fmt.Println("Fredag är veckodag")
	// case "Saturday":
	// 	fmt.Println("Lördag är helgen")
	// case "Sunday":
	// 	fmt.Println("Söndag är helgen")

	// }

	// type Switch case using function:

	// WhoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("it's an integer data type i.e. ", t)
	// 	case string:
	// 		fmt.Println("It's a string data type i.e. ", t)
	// 	case bool:
	// 		fmt.Println("It's a boolean data type i.e. ", t)
	// 	default:

	// 		fmt.Println("Not a datatype", t)

	// 	}
	// }
	// WhoAmI("Himanshu")

	// imported time function

	switch time.Now().Weekday() {
	case time.Monday, time.Friday, time.Thursday:
		fmt.Println("These are weekdays/Veckodag")
	case time.Saturday, time.Sunday:
		fmt.Println("These are Weekend/Helgedag")
	default:
		fmt.Println("Not a valid day")
	}

}
