package main

import (
	"books/cmd/handlers"
	"books/cmd/repositories"
	"books/cmd/routes"
	"log"
	"net/http"
)

func main() {
	repo := repositories.NewInMemoryBookRepository()
	bookHandler := handlers.NewBookHandler(repo)
	routes.SetupBookRoutes(bookHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
