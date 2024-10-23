/*
	Using Pointers & Passing Pointers To Functions


*/

package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	// 2.
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)

	fmt.Println(age) // 32
}

// 1. need to change the type of param to pointer -> *int
func getAdultYears(age *int) int {
	// now, we need to use the pointer here as well -> *age
	return *age - 18
}
