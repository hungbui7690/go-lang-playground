/*
	make function


*/

package main

import "fmt"

func main() {
	// userNames := []string{} // empty slice
	userNames := make([]string, 2, 5) // create a slice with length 2 and capacity 5

	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames) // [Julie _ Max Manuel]
}
