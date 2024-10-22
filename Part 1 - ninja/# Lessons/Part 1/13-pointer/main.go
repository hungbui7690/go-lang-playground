/*
  Pointer
	- to get address of a value: &name
	- to pass a pointer: *name


*/

package main

import "fmt"

// func updateName(n string) {
// 	n = "wedge"
// }

// pass by reference -> pass the pointer
func updateName(n *string) {
	*n = "wedge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateName(m) // using pointer as arg, can pass &name directly instead of "m" as well 
	fmt.Println(name) // wedge -> we can see that the original value is changed from "tifa" to "wedge"

}

/*

|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "tifa"  | p0x001  |
|---------|---------|

*/
