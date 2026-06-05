package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fabio98p/homelabinfra_be/internal/db"
	"github.com/fabio98p/homelabinfra_be/internal/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	h := handlers.NewHandler(database)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", h.GetTodos)
	mux.HandleFunc("POST /todos", h.CreateTodo)
	mux.HandleFunc("GET /todos/{id}", h.GetTodo)
	mux.HandleFunc("PUT /todos/{id}", h.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", h.DeleteTodo)
	mux.HandleFunc("OPTIONS /todos", h.Options)
	mux.HandleFunc("OPTIONS /todos/{id}", h.Options)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}
