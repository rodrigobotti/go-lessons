package main

import (
	"fmt"
)

// Property struct holding property data
type Property struct {
	X     int
	Y     int
	Name  string
	Value float64
}

func main() {
	home := new(Property)
	fmt.Printf("Home: %[1]p - %+[1]v\n", home)
	countryHome := Property{X: 10, Y: 10, Name: "Country Home", Value: 100}
	fmt.Printf("CountryHome: %p - %+v\n", &countryHome, countryHome)
	changeProperty(&countryHome)
	fmt.Printf("CountryHome: %p - %+v\n", &countryHome, countryHome)
}

func changeProperty(property *Property) {
	property.X = property.X + 10
	property.Y = property.Y - 5
	property.Name = "* " + property.Name + " *"
}
