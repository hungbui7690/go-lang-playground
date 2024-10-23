/*
	Struct Pointers


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

	var appUser user = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// 2. pass the pointer of the appUser
	outputUserDetails(&appUser)
}

// 1. just to prevent copying
func outputUserDetails(u *user) {
	// 2 methods below return the same result 
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	// fmt.Println((*u).firstName, (*u).lastName, (*u).birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
