/*
	Instantiating Structs & Struct Literal Notation


*/

package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// 1. Method 1
	// var appUser user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(), // if missing any of these fields -> that value will be zero value
	// }

	// 2. Method 2: omit the key -> order must be the same as we defined above
	appUser := user{
		userFirstName,
		userLastName,
		userBirthdate,
		time.Now(),
	}
	fmt.Println(appUser) // this will return the address

	// 4. passing struct as a argument
	outputUserDetails(appUser)
}

// 3.
func outputUserDetails(u user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
