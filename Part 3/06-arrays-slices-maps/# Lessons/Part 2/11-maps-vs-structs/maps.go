/*
	Maps vs Structs
	- structs -> cannot be modified -> more complex
		-> used to describe an entity
	- maps -> simpler


*/

package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

// DON'T DO THIS
type Websites struct{
	goggle string
	aws string
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
}
