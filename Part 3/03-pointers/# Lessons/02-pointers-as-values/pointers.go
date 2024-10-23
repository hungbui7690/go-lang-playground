/*
	Create Pointers
	- * -> pointer
	- & -> address


*/

package main

import "fmt"

func main() {
	age := 32 

	// # Create pointer
	var agePointer *int = &age
	// agePointer := &age 

	fmt.Println("*agePointer:", *agePointer) // value
	fmt.Println("agePointer:", agePointer) // address
}
