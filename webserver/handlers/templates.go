package handlers

import "html/template"

var (
	// HelloTemplate holds parsed 'views/hello.html' Template reference
	HelloTemplate = template.Must(template.ParseFiles("webserver/views/hello.html"))
)
