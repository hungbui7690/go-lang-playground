/*
	Working w Arrays
	- {} = array
	- [4]string{} = array of 4 strings
	- array = fixed size


*/

package main

import "fmt"

// # for complex types
// type Product struct{
// 	id    string
// 	title string
// 	price float64
// }

// type TemperatureData struct{
// 	day1 float64
// 	day2 float64
// 	day3 float64
// 	day4 float64
// 	day5 float64
// }

func main() {
	// var productNames [4]string = [4]string{"A Book"}
	var productNames = [4]string{"A Book"} // 4 elements
	productNames[2] = "A Carpet"
	fmt.Println(productNames)
	
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[2])
}
