/*
	Using Third-Party Packages
	- https://pkg.go.dev/ -> search for packages
	- Example: https://github.com/Pallinder/go-randomdata


	++++++++++++++
	Installation
	- go get github.com/Pallinder/go-randomdata
		-> check .mod file

	- later, when someone download our project from github
		~~ go get


*/

package main

import (
	"fmt"

	"example.com/app/fileops"
	"github.com/Pallinder/go-randomdata" // 1.
)

const accountBalanceFile = "balance.txt"

func main() {
	fmt.Println(randomdata.SillyName()) // 2.

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	presentOptions()
	for {

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)


		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			accountBalance += depositAmount 
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			}

			accountBalance -= withdrawalAmount 
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank")
			return
			// break
		}
	}
}
