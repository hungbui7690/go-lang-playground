/*
	Format Strings
	- Sprintf() -> save format strings
	- Print() -> print
	- Println() -> print with new line
	- Printf() -> print format strings
		# https://pkg.go.dev/fmt
		# %v
		# %#v
		# %%: literal percent sign
		# %f: float
		# %d: int -> base10
		# %s: string


*/

package main

import (
	"fmt"
)

func main() {
	futureValue := 1234.567
	futureRealValue := 1111.4321

	fmt.Printf("Future Value: %-25v\n", futureValue)
	fmt.Printf("Future Real Value: %-25v\n", futureRealValue)
	fmt.Printf("Future Value: %.1f\tFuture Real Value: %.2f\n", futureValue, futureRealValue)

	// save formatted string to variable
	formattedFV := fmt.Sprintf("Decimal: %.1f\n", 5.555) 
	fmt.Println(formattedFV)
}
