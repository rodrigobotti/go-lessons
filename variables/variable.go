package main

import (
	"fmt"
)

var (
	address, phone string  // address = ""
	qty            int     // qty = 0
	bought         bool    // bought = false
	val            float64 // val = 0.00
	words          rune
	// Public variable (starts with uppercase)
	Public string
)

func main() {
	test := "Testing"
	fmt.Printf("Testing: %s\n", test)
	fmt.Printf("address: %s\n", address)
	fmt.Printf("qty: %d\n", qty)
	fmt.Printf("bought: %t\n", bought)
	fmt.Printf("words: %v\n", words)
	fmt.Printf("words: %.2f\n", val)
}
