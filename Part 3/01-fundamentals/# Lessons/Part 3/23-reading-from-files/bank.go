/*
	Reading from Files


*/

package main

import (
	"fmt"
	"os"
	"strconv" // # to convert from string to float
)


const accountBalanceFile = "balance.txt" // 1.

// 2.
func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile) // data is of type []byte
	fmt.Print(data)
	balanceText := string(data) // convert from []byte to string
	balance, _ := strconv.ParseFloat(balanceText, 64) // convert from string to float
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile() // 3.

	fmt.Println("Welcome to Go Bank!")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int

	for {
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
				writeBalanceToFile(accountBalance)
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
				writeBalanceToFile(accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for choosing our bank")
				return
				// break
			}
	}
}
