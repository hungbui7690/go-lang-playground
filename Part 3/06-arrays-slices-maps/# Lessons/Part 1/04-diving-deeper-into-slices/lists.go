/*
	Diving Deeper Into Slices
	- A slice has both a length and a capacity.
	- The length of a slice is the number of elements it contains.
	- The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
		-> cap === capacity


*/

package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	
	
	featuredPrices := prices[1:] // has 3 elements
	highlightedPrices := featuredPrices[:1] // has 1 element
	fmt.Println("highlightedPrices", highlightedPrices)
	fmt.Println("prices", prices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1, 3 -> because the underlying array (featuredPrices) has 3 elements -> capacity = 3


	highlightedPrices = highlightedPrices[:3]	// has 3 elements
	fmt.Println(highlightedPrices) // [10.99 9.99 45.99]
	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 3, 3
}
