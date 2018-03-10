package main

import (
	"fmt"
	"strings"

	"github.com/rodrigobotti/go-lessons/packages/operator"
	"github.com/rodrigobotti/go-lessons/packages/prefix"
)

var (
	// UserName system's user's name
	UserName  = "Botti"
	separator = strings.Repeat("*", 20)
)

func aroundSeparator(fn func()) func() {
	return func() {
		fmt.Println(separator)
		fn()
		fmt.Println(separator)
	}
}

var printState = aroundSeparator(func() {
	fmt.Printf("UserName: %s\n", UserName)
	fmt.Printf("Capitol prefix is: %d\n", prefix.Capitol)
	fmt.Printf("Operator is: %s\n", operator.OperatorName)
	fmt.Printf("Operator prefix is: %s\n", operator.OperatorPrefix)
})

func main() {
	printState()
}
