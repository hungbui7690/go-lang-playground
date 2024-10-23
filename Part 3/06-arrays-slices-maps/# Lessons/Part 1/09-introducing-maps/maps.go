/*
	Introducing Maps
	- maps are key/value pairs


*/

package main

import "fmt"

func main() {
	// key=string, value=string
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Amazon Web Services"])
}
