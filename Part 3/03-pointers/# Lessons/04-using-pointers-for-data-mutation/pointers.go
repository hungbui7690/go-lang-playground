/*
  Using Pointers For Data Mutation -309K


*/

package main

import "fmt"

func main() {
	age := 32 
	var agePointer *int = &age

	editAgeToAdultYears(agePointer) 
	fmt.Println(age) // 14
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // mutate the value
}
