/*
  Printing & Formatting Strings
	- Print
	- Println
	- Printf: print w format
		# %_ = format specifier
		# %v = value in default format
		# %f = float
		# %0.2f = float with 2 decimal points
		# %q = quotes
		# %T = type
		# %s = string
		# %d = int
		# %b = binary
		# %t = true/false
	- Sprintf: save formatted strings


*/

package main

import "fmt"

func main() {
	age := 35
	name := "mario"

	// ------------
	// Print
	// ------------
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")


	// ------------
	// Println
	// ------------
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")

	// # with variables
	fmt.Println("my age is", age, "and my name is", name) 


	// ------------
	// # Printf -> format variables
	// ------------
	fmt.Printf("my age is %v and my name is %v \n", age, name) 
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age) // # = number
	fmt.Printf("you scored %f points! \n", 225.55) // 225.550000           
	fmt.Printf("you scored %0.1f points! \n", 225.55) // 225.6
	fmt.Printf("you scored %0.2f points! \n", 225.55) // 225.55


	// ------------	
	// Sprintf 
	// ------------	
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name) // save formatted string into "str"
	fmt.Println("the saved string is:", str)


	// see more format specifiers here:
	// https://golang.org/pkg/fmt/

}