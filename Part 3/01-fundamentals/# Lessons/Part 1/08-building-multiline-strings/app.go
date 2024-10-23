/*
	Building Multiline Strings


*/

package main

import (
	"fmt"
)

func main() {
	futureValue := 1234.567
	futureRealValue := 1111.4321
	
	// use backtick for multiple lines -> no ned
	fmt.Printf(`
		Future Value: %.1f
		Future Real Value: %.2f
	`, 
	futureValue, futureRealValue)
	
}
