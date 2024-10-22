/*
  Maps
	- similar to objects


*/

package main

import "fmt"

func main() {
	// k=string : v=float64
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// k=int : v=string
	phoneBook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[845775485])

	phoneBook[984759373] = "bowser"
	fmt.Println(phoneBook)

	phoneBook[647583927] = "yoshi"
	fmt.Println(phoneBook)
}