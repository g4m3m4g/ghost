package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	r.Post("/create/{roomId}", createRoom)
	r.Get("/ws/{roomId}", wsHandler)

	go cleanupRooms()

	http.ListenAndServe(":"+port, r)
}