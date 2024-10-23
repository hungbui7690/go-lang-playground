/*
	User Input


*/

package main

import (
	"fmt"
)

func main() {
	var investmentAmt float64 = 1000 // initial value


	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmt) // get user input + update investmentAmt

	fmt.Println(investmentAmt)

}
