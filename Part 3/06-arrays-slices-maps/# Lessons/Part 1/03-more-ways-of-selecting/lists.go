/*
	More Ways Of Selecting Slices


*/

package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	
	// more ways of selecting parts of an array
	featuredPrices := prices[1:] // from index 1 to the end
	highlightedPrices := featuredPrices[:1] // return only the first element
	fmt.Println(highlightedPrices)
}
