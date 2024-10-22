/*
  Structs vs Custom Types
	- we want to create bill type that has name, items, and tip
		-> custom type
		-> struct


*/

package main

import "fmt"

// 1. custom type "bill" -> struct
type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// 2. make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		// items: map[string]float64{"pie": 5.99, "cake": 3.99}
		tip:   0,
	}
	return b
}

func main() {
	// 3.
	mybill := newBill("mario's bill")
	fmt.Println(mybill)
}