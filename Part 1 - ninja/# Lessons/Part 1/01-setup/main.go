/*
  Intro
  - fast, statically typed, compiled
    -> no need to interpret line by line
  - general purpose language
    -> use in many kinds of apps
  - built-in testing support
  - oop -> in it's own way


\\\\\\\\\\\\\\\\\\\\\\\\

  Setup
  - https://go.dev/dl/
    -> download
  - extension: go


\\\\\\\\\\\\\\\\\\\\\\\\

	Run
	- go run main.go


*/

package main // tell go that it should compile to a standalone executable file at the end

import "fmt" // std lib ->  formatted I/O -> print message to console

// only 1 main function in our app
func main() {
	fmt.Println("Hello, playground")
}