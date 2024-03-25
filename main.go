package main

import (
	"net/http"
	"fmt"
	"html/template"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("RUN PROGRAM on Port 1313 ----------")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		tmpl, _ := template.New("").ParseFiles("templates/index.html")
		tmpl.ExecuteTemplate(w, "base", nil)
	})
	http.ListenAndServe(":1313", r)
}
