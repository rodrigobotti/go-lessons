package main

import (
	"fmt"
)

func main() {
	months := 0
	debt := true
	city := "test"

	// < > != == <= >= && ||
	if months <= 6 {
		fmt.Println("Short-term debt")
	}
	if debt {
		fmt.Println("Is in debt")
	}
	if !debt {
		fmt.Println("Not in debt")
	}
	if city == "test" {
		fmt.Println("Lives in state 'test'")
	}
	if description, status := debtForHowLong(months); status {
		fmt.Println("Client's situation is:", description)
		return
	}
	// fmt.Println("Status is: ", descritpion, status) -> variables only available inside if block
	fmt.Println("Thank you for choosing us")
}

func debtForHowLong(months int) (description string, status bool) {
	if months > 0 {
		status = true
		description = "Is in debt"
		return
	}
	description = "Not in debt"
	return
}
