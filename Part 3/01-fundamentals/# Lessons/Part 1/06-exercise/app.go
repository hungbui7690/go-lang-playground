/*
	Exercise - Profit Calculator
	- Ask for revenue, expenses, and tax rate
	- Calculate Earning Before Tax (Ebt) and Earning After Tax (Profit)
	- Calculate ratio between Ebt and Profit (Ebt/Profit)
	- Output Ebt, Profit, and ratio


*/

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}
