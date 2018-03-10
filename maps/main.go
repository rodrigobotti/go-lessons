package main

import (
	"fmt"

	"github.com/rodrigobotti/go-lessons/maps/model"
)

var cache map[string]model.Property

func main() {
	const key = "Yellow House"
	house := model.Property{Name: key, X: 1, Y: 2}
	house.SetValue(123123)
	cache = make(map[string]model.Property, 0)
	cache[key] = house
	if property, found := cache[key]; found {
		fmt.Printf("Found %s with value %+v\n", key, property)
	}

	aptmt := model.Property{Name: "Apartment", X: 10, Y: 5}
	aptmt.SetValue(400000.00)

	cache[aptmt.Name] = aptmt
	fmt.Println("Cache has size of", len(cache))

	for key, val := range cache {
		fmt.Printf("[%s] : %+v\n", key, val)
	}

	if _, found := cache[key]; found {
		delete(cache, key)
	}
	delete(cache, "BLAH")
	fmt.Println("Cache has size of", len(cache))
}
