/*
	Struct Embedding


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

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 3.
	admin := user.NewAdmin("test@example.com", "test123")

	// 4. admin still can run the same methods as appUser (step 1)
	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

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
