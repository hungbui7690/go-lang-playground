/*
	Mutating Maps


*/

package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com" // add
	fmt.Println(websites)

	delete(websites, "Google") // delete
	fmt.Println(websites)
}
