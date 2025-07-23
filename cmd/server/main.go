package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce-app/internal/database"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Connect to DB and run migrations
	database.ConnectAndMigrate()

	r := chi.NewRouter()

	// Simple health check route
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
