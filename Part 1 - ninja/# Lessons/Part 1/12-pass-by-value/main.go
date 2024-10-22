/*
  Pass By Value
	- Go makes "copies" of values when passes into functions
	- Group A: strings, ints, bools, floats, arrays, structs
		Group B: slices, maps, functions


*/

package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

// # pass by value
func updateName(x string) string {
	x = "wedge"
	return x
}

// # pass by reference -> pass the pointer
func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types 
	// non-pointer wrapper values
	name := "tifa"
	name = updateName(name)
	fmt.Println(name) // wedge

	// group B types 
	// pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}
	updateMenu(menu) 
	fmt.Println(menu) // map[coffee:2.99 ice cream:3.99 pie:5.95] -> we changed the value in the map
}