package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rodrigobotti/go-lessons/webserver/models"
	"github.com/rodrigobotti/go-lessons/webserver/utils"
)

// Greet serializing 'Hello World' response
func Greet(w http.ResponseWriter, r *http.Request) {
	name, found := utils.QueryParam("name", r)
	if !found {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

// GreetView renders the hello view
func GreetView(w http.ResponseWriter, r *http.Request) {
	page := models.Page{Hour: time.Now().Format("15:04:05")}
	if err := HelloTemplate.Execute(w, page); err != nil {
		fmt.Println("[GreetView] error rendering template", err.Error())
		http.Error(w, "Error when rendering template", http.StatusInternalServerError)
	}
}
