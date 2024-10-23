/*
	Splitting Code Across Multiple Packages
	1. fileops/fileops.go

		🌲 import module/package
		🌿 package.fn()


\\\\\\\\\\\\\\\\\\\\\\\\\\\

	Rules
	- functions with the name starts with capital letter are exported
		-> fmt.Println()
		-> fmt.Print()
		-> fmt.Sprint()
		-> fmt.Scan()
		...


*/

package main

import (
	"fmt"

	"example.com/app/fileops" // 1. import module/package
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile) // 2.
	

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("---------")
		// panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	presentOptions()
	for {
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// wantsCheckBalance := choice == 1

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

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
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

			accountBalance -= withdrawalAmount // accountBalance = accountBalance + depositAmount
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
