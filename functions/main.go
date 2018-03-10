package main

import (
	"fmt"

	"github.com/rodrigobotti/go-lessons/functions/math"
)

func main() {
	res := math.Multiply(2, 4)
	fmt.Printf("[m] res: %d\n", res)
	res = math.Add(2, 1)
	fmt.Printf("[a] res: %d\n", res)
	res = math.Calculate(math.Add, 4, 5)
	fmt.Printf("[c][a] res: %d\n", res)
	res = math.Calculate(func(x int, y int) int { return x - y }, 5, 3)
	fmt.Printf("[c][l] res: %d\n", res)
	res = math.Div(20, 2)
	fmt.Printf("[d] res: %d\n", res)
	res, rem := math.DivWithRemainder(7, 2)
	fmt.Printf("[dr] res: %d, rem: %d\n", res, rem)
}
