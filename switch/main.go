package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := 3
	// breaking is not necessary
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch runtime.GOOS {
	case "windows":
		fallthrough
	case "darwin":
		fmt.Println("MACOS")
	default:
		fmt.Println("Whatevs...")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work!")
	}

	num = 10
	switch {
	case num < 10:
		fmt.Println("Yup")
	case num >= 10 && num < 100:
		fmt.Println("Two digits")
	case num >= 100:
		fmt.Println("Number is too big")
	}

}
