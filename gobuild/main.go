package main

import (
	"encoding/json"
	"fmt"

	"github.com/rodrigobotti/go-lessons/gobuild/model"
)

func main() {
	home := model.Property{}
	fmt.Printf("home: %+v\n", home)
	aptmt := model.Property{X: 17, Y: 56, Name: "My apartment"}
	aptmt.SetValue(400000.00)
	fmt.Printf("Apartment: %+v\n", aptmt)
	home.SetValue(123123)
	fmt.Printf("Home value: %.2f\n", home.GetValue())
	objJSON, _ := json.Marshal(aptmt)
	fmt.Println("Apartment: ", string(objJSON))
}
