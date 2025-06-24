package main

import (
	"github.com/a-h/templ"
	"net/http"
	"willoh/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Parent()))
	http.ListenAndServe(":8080", nil)
}
