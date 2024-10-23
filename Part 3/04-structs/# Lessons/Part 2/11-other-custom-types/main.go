package main

import "fmt"

// custom type
type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var name customString = "Joe Doe"

	name.log()
}
