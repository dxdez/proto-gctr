package main

import (
	"net/http"
	"fmt"
	"html/template"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Indicate that the application is running
	fmt.Println("RUN PROGRAM on Port 8080 ----------")

	// Start database
	dbRunOrError := openDB()
	if dbRunOrError != nil {
		log.Panic(dbRunOrError)
	}
	defer closeDB()
	dbRunOrError = setupDB()
	if dbRunOrError != nil {
		log.Panic(dbRunOrError)
	}

	// Parse Templates
	parseOrError = parseTemplates()
	if parseOrError != nil {
		log.Panic(parseOrError)
	}


	// Establish router
	appRouter := chi.NewRouter()
	appRouter.Use(middleware.Logger)
	appRouter.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		currentTemplate.ExecuteTemplate(w, "base", nil)
	})

	// Run application
	http.ListenAndServe(":8080", appRouter)
}
