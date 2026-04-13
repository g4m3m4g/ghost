package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Post("/create/{roomId}", createRoom)
	r.Get("/ws/{roomId}", wsHandler)

	go cleanupRooms()

	http.ListenAndServe(":8080", r)
}