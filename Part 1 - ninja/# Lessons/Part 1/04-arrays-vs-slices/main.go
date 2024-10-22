/*
  Arrays & Slices
	- in go, arrays is {}, not []
	- array: [3]int -> fixed length
	- slice: []int



*/

package main

import "fmt"

// ------------------
// arrays
// ------------------
// func main() {
// 	// var ages [3]int = [3]int{20, 25, 30}
// 	var ages = [3]int{20, 25, 30} // type annotation on the left side can be ignored

// 	names := [4]string{"yoshi", "mario", "peach", "bowser"}
// 	names[1] = "luigi"

// 	fmt.Println(ages, len(ages))
// 	fmt.Println(names, len(names))
// }

// ------------------
// slices (use arrays under the hood)
// ------------------
func main() {
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85) // return new slice
	fmt.Println(scores, len(scores))

	// slice ranges -> slice the original array
	names := [4]string{"yoshi", "mario", "peach", "bowser"}	
	rangeOne := names[1:4] // [mario peach bowser] -> doesn't include pos 4 element
	rangeTwo := names[2:]  // [peach bowser] -> includes the last element
	rangeThree := names[:3] // [yoshi mario peach]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo) 
	fmt.Println(rangeThree) 
	fmt.Printf("the type of rangeOne is %T \n", rangeOne) // the type of rangeOne is []string 

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne) // [mario peach bowser koopa]

}