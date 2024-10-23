/*
	Type Annotations


*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmt float64 = 1000 // explicit type assignment
	const expectedReturn = 5.5 // constant
	years := 10 // type inference

	// var investmentAmt, years float64 = 1000, 10
	// expectedReturn := 5.5

	futureValue := investmentAmt * math.Pow(1 + expectedReturn / 100, float64(years))
	fmt.Println(futureValue)

	var name, age = "Joe Doe", 29
	fmt.Println(name, age)
}
