package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i is", i)
	}

	val, test := 0, true
	for test {
		val++
		if val%5 == 0 {
			test = false
		}
		fmt.Println("value is", val)
	}
	for {
		val--
		fmt.Println("value is", val)
		if val == 0 {
			break
		}
	}
	for i, letter := range "I love writing code in go" {
		fmt.Printf("Text[%d] = %q\n", i, letter)
	}
}
