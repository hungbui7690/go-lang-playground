/*
	Exposing Methods & A Different Constructor Function Name
	1. create user/user.go file
		-> place the user struct code there
		-> rename the exposed methods
	2. import

	🌿 rename the struct + apply in structs.go as well


*/

package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User // 2. change this as well

	// 3. if we want to access firstName here -> need to rename to FirstName as well
	// appUser = &user.User{
	// 	firstName: userFirstName,
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
