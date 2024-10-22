/*
  Variables & String & Int


*/

package main

import "fmt"

func main() {

	// ------------------
	// string variables
	// ------------------

	var nameOne string = "mario"
	var nameTwo = "luigi" // inferred as string
	var nameThree string
	nameFour := "yoshi" // another way to initialize a variable -> shorthand -> allowed inside functions only
	fmt.Println(nameFour) // yoshi

	fmt.Println(nameOne, nameTwo, nameThree) // mario luigi 

	nameOne = "peach"
	nameThree = "bowser"

	fmt.Println(nameOne, nameTwo, nameThree) // peach luigi bowser


	// ------------------
	// int variables
	// ------------------

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)


	// ------------------
	// bits & memory
	// ------------------

	// var numOne int8 = 25 // 8-bits -> int8: -128 to 127 -> int16 int32 int64 
	// var numTwo int8 = 128 // too large a number for 8-bit
	// var numTwo uint = -25 unsigned ints cannot be negative -> 0-255


	// ------------------
	// numbers
	// ------------------

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// for more info see https://golang.org/ref/spec#Numeric_types

}