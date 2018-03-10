package main

import (
	"fmt"

	"github.com/rodrigobotti/go-lessons/interfaces/model"
)

func main() {
	jojo := model.Bird{Name: "Jojo"}
	wakeUpWithCluck(jojo)
	hearBirdInALake(jojo)
}

func wakeUpWithCluck(chicken model.Chicken) {
	fmt.Println(chicken.Cluck())
}

func hearBirdInALake(duck model.Duck) {
	fmt.Println(duck.Quack())
}
