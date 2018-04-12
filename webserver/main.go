package main

import (
	"fmt"
	"net/http"

	"github.com/rodrigobotti/go-lessons/webserver/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Greet)
	http.HandleFunc("/view/hello", handlers.GreetView)
	fmt.Println("listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
