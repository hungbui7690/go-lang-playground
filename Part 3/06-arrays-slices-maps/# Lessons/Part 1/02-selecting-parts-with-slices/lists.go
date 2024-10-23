/*
	Selecting Parts of Arrays With Slices


*/

package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	
	// slice -> select parts of an array -> return a slice
	featuredPrices := prices[1:3] // return index 1 and 2
	fmt.Println(featuredPrices)
}
