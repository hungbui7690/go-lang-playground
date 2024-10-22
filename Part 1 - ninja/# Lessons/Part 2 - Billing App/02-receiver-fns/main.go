/*
  Receiver Functions
	- func (b bill) format() string {}
		-> receiver function
	- %-25v = 25 characters long



*/

package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	fmt.Println(mybill.format()) 
	// Bill breakdown:
	// pie:                      ...$5.99
	// cake:                     ...$3.99
	// total:                    ...$9.98
}