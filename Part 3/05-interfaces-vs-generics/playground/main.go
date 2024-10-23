/*
	Interfaces, Dynamic Types & Limitations
	Generics


*/

package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)

	// resultX := add(1.5, 2.5)
	// fmt.Println(resultX)

	// resultZ := add('x', 'y') // not work
	// fmt.Println(resultZ)
}


// 1. limit of interfaces
// func add(a, b interface{}) {
// 	return a + b
// }

// 2. solution for step 1 -> but require a lot of work
// func add(a, b interface{}) interface{} {
// 	aInt, aIsInt := a.(int)
// 	bInt, bIsInt := b.(int)

// 	if aIsInt && bIsInt {
// 		return aInt + bInt
// 	}

// 	aFloat, aIsFloat := a.(float64)
// 	bFloat, bIsFloat := b.(float64)

// 	if aIsFloat && bIsFloat {
// 		return int(aFloat) + int(bFloat)
// 	}

// 	aString, aIsString := a.(string)
// 	bString, bIsString := b.(string)	

// 	if aIsString && bIsString {
// 		return aString + bString
// 	}

// 	return nil
// }

// 3. generics solution 1 -> not work 
// func add[T any](a, b T) T {
// 	return a + b
// }

// 4. T can be int, float64 or string
func add[T int | float64 | string](a, b T) T {
	return a + b
}

