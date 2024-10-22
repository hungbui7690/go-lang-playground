/*
  Package Scope
	- same package name
		-> no need to import

	🚀 go run main.go greetings.go


*/

package main

import "fmt"

var score = 99.5

// cannot use shorthand outside of functions
// scoreTwo := 50

func main() {
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}