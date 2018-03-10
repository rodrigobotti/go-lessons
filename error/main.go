package main

import (
	"encoding/json"
	"fmt"

	"github.com/rodrigobotti/go-lessons/error/model"
)

func main() {
	home := model.Property{Name: "Home", X: 1, Y: 2}
	if err := home.SetValue(50000000); err != nil {
		fmt.Println("[main] Error when setting property value:", err.Error())
		if err == model.ErrHighPropertyValue {
			fmt.Println("Realter, please check the value rating on this property")
		}
	}

	if _, err := json.Marshal(home); err != nil {
		fmt.Println("[main] Error when serializing JSON:", err.Error())
		return
	}
}
