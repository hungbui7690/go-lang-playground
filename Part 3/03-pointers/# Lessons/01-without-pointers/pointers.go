package main

import "fmt"

func main() {
	age := 32 // regular variable

	fmt.Println("Age:", age)

	// param will be copied -> waste memory
	adultYears := getAdultYears(age)
	fmt.Println(adultYears)
}


func getAdultYears(age int) int {
	return age - 18
}
