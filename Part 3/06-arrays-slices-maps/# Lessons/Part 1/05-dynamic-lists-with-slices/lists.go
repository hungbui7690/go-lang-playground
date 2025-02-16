/*
	Building Dynamic Lists With Slices
	- slice = a dynamically sized array


*/

package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(cap(prices), len(prices)) // 2, 2
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)
}
