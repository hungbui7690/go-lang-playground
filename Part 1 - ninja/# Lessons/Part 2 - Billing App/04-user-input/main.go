/*
  User Input



*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 2. hover on scanner/reader variable to see its type
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

// 3.
func createBill() bill {
	scanner := bufio.NewReader(os.Stdin) // scanner or reader

	// 1. Method 1
	// fmt.Print("Create a new bill name: ")
	// name, _ := scanner.ReadString('\n') // name, _ === err
	// name = strings.TrimSpace(name)


	// Method 2: create function getInput and use here
	name, _ := getInput("Create a new bill name: ", scanner)

	b := newBill(name)
	fmt.Println("Bill Name - ", b.name)

	return b
}

// 4.
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	fmt.Println("Option - ", opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}