package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99}, // 1.
		tip:   0,
	}
	return b
}

// 2. we want to associate format() with bill object -> func (b bill) format() string{} -> with this setup, we can use bill.format()
func (b bill) format() string { 
	fs := "Bill breakdown:\n" // fs = format string
	var total float64 = 0

	// k: pie, v: 5.99 -> we want to add colon after k -> k+":"
	for k, v := range b.items {
		// fs += fmt.Sprintf("%v ...$%v\n", k+":", v)
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}