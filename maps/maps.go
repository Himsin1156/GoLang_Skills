package main

import (
	"fmt"
	"maps"
)

// Map is like -> hash , object, dictionary

func main() {

	// 1. Creating map

	map1 := make(map[string]string) // make(map[key]value)  here key is string and value needed also string hence used string for both
	map2 := make(map[string]int)    //   make(map[key]value)  here key is string i.e mmob number but needed number as integer so value is int. and key is string.

	// filling elements  or setting up the elements

	map1["name"] = "Himanshu"
	map1["Mobile"] = "Iphone"
	map1["Email"] = "himanshu1156@gmail.com"
	map1["Address"] = "Stockholm"
	map2["MobileNumber"] = 761593747

	// get an element or printing the output:

	//fmt.Println(map1["name"], map1["Email"], map1["Mobile"], map1["Address"], map2["MobileNumber"])

	/*******/
	// 2. Delete an element from map
	/*******/

	//delete(map1, "Email")

	/*******/
	// 3. Clear all Elements from map
	/*******/

	//clear(map1)

	//fmt.Println(map1, map2)

	/***********************/
	// 4. Another way to create map without using make function
	/***********************/

	//geomap := map[string]string{"Name": "Himanshu", "Address": "Kanpur", "Office": "Netinsight_AB"}
	geomap2 := map[string]int{"Age": 30, "Mobile": 761593747}
	geomap3 := map[string]int{"Age": 30, "Mobile": 761593747}
	// G, ok := geomap["Name"]
	// fmt.Println(G)
	// if ok {
	// 	fmt.Println("Key is ok")
	// } else {
	// 	fmt.Println("Key is not-ok")
	// }

	fmt.Println(maps.Equal(geomap3, geomap2))

	// fmt.Println(geomap["Name"])
	// fmt.Println(geomap2["Mobile"])
}
