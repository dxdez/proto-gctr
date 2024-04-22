package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Indicate that the application is running
	fmt.Println("RUN PROGRAM on Port 8080 ----------")

	// Start database
	runOrError := openDB()
	if runOrError != nil {
		log.Panic(runOrError)
	}
	defer closeDB()
	runOrError = setupDB()
	if runOrError != nil {
		log.Panic(runOrError)
	}

	// Parse Templates
	runOrError = parseTemplates()
	if runOrError != nil {
		log.Panic(runOrError)
	}

	// Establish router
	appRouter := chi.NewRouter()
	appRouter.Use(middleware.Logger)
	appRouter.Get("/", handleGetItems)
	appRouter.Post("/items", handleCreateItem)
	appRouter.Put("/items/{id}/toggle", handleToggleItem)

	// Run application
	http.ListenAndServe(":8080", appRouter)
}
