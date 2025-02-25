package main

import (
	"github.com/a-h/templ"
	"net/http"
	"willoh/views"
)

func main() {
	// http.Handle("/", templ.Handler(views.dog("willoh")))
	// http.ListenAndServe(":8080", nil)
	views.Bar()
}
