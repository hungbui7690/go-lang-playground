/*
  The Standard Library
	- https://pkg.go.dev/std


\\\\\\\\\\\\\\\\\\\\\\\\\

	- strings.Contains()
	- strings.ReplaceAll()
	- strings.ToUpper()
	- strings.Index()
	- strings.Split()


\\\\\\\\\\\\\\\\\\\\\\\\\

	- sort.Ints()
	- sort.Strings()
	- sort.SearchInts(ages, 30)


*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// --------------------
	// # "strings" library
	// --------------------
	fmt.Println(strings.Contains(greeting, "hello")) // true
	fmt.Println(strings.Contains(greeting, "bye"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	// slice of int
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}


	// --------------------
	// # sort library
	// --------------------
	sort.Ints(ages) // mutating the original slice
	fmt.Println(ages) // [20 25 30 35 45 50 60 75]

	index := sort.SearchInts(ages, 30) // return index
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names) // [bowser luigi mario peach yoshi]

	fmt.Println(sort.SearchStrings(names, "bowser")) // 0

}