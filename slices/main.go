package main

import (
	"fmt"
	"strings"
)

func printStringSlice(slice []string) {
	fmt.Println(slice, len(slice), cap(slice))
}

func printIntSlice(slice []int) {
	fmt.Println(slice, len(slice), cap(slice))
}

func removeItemAt(index int, slice []string) []string {
	tempHead := slice[:index]            // head not including item
	tempTail := slice[index+1:]          // tail not including item
	return append(tempHead, tempTail...) // new slice with item removed
}

// cannot be used with []string, []int, []<some_struct> or []*<some_struct>
// leaving it here just for reference
func printSlice(slice []interface{}) {
	fmt.Println(slice, len(slice), cap(slice))
	for index, item := range slice {
		fmt.Printf("[%d] = %+v", index, item)
	}
}

func main() {
	var nums []int
	printIntSlice(nums)
	nums = make([]int, 5)
	printIntSlice(nums)
	nums = append(nums, 99)
	printIntSlice(nums)

	capitols := []string{"Lisboa"}
	printStringSlice(capitols)
	capitols = append(capitols, "Brasilia")
	printStringSlice(capitols)
	capitols[1] = "Brasília"
	printStringSlice(capitols)

	cities := make([]string, 5)
	cities[0] = "New York"
	cities[1] = "London"
	cities[2] = "Madeira"
	cities[3] = "Tokyo"
	cities[4] = "Singapore"
	// cities[5] = "Blah" --> panic: runtime error: index out of range
	printStringSlice(cities)
	for index, city := range cities {
		fmt.Printf("City[%d] = %s\n", index, city)
	}

	fmt.Println(strings.Repeat("*", 10), "Slices advanced", strings.Repeat("*", 10))

	asianCapitols := cities[3:5] // last value is non-inclusive
	printStringSlice(asianCapitols)

	head2 := cities[:2]
	printStringSlice(head2)

	tail2 := cities[len(cities)-2:]
	printStringSlice(tail2)

	temp := removeItemAt(2, cities)
	copy(cities, temp)
	printStringSlice(cities)
}
