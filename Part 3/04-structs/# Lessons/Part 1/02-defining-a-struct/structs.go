/*
	Structs
	- define custom types


*/

package main

import (
	"fmt"
	"time"
)

// # define a struct
type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time // #
}

// require more work for this setup
func outputUserDetails(firstName, lastName, birthdate string) {
	// ...
	fmt.Println(firstName, lastName, birthdate)
}

func main() {
	outputUserDetails("John", "Doe", "01/01/2000")
}


