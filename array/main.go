package main

import (
	"fmt"
)

func main() {
	var test [3]int
	test[0] = 3
	test[1] = 2
	test[2] = 1
	fmt.Println("Array size is", len(test))

	singers := [2]string{"Michael Jackson", "John Lennon"}
	fmt.Printf("Singers are %s\n", singers)

	capitols := [...]string{"Lisboa", "Brasilia", "Luanda", "Maputo"}
	fmt.Println("Array size is", len(capitols))
	for index, city := range capitols {
		fmt.Printf("capitol[%d] = %s\n", index, city)
	}
}
