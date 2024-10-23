/*
	An Alternative Return Value Syntax


*/

package main

import "fmt"

func main() {
	fmt.Println(calculateFutureValues())
}

// add return variables here
func calculateFutureValues() (dec1 float64, dec2 float64) {
	dec1 = 0.5
	dec2 = 0.5

	return dec1, dec2
}