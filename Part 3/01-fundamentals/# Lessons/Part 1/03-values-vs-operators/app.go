/*
	Values vs Operators


*/

package main

import "fmt"

func main() {
	var investmentAmt = 1000
	var expectedReturn = 5.5

	// error -> mismatch type
	// var futureValue = investmentAmt * ( 1 + expectedReturn / 100)  

	// type conversion
	var futureValue = float64(investmentAmt) * ( 1 + expectedReturn / 100)
	fmt.Println(futureValue)
}
