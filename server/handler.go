package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

var (
	rooms = make(map[string]*Room)
	mu    sync.Mutex

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func createRoom(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")

	mu.Lock()
	defer mu.Unlock()

	if _, exists := rooms[roomId]; exists {
		http.Error(w, "room exists", 400)
		return
	}

	room := &Room{
		id:        roomId,
		clients:   make(map[*Client]bool),
		broadcast: make(chan []byte),
		createdAt: time.Now(),
	}

	rooms[roomId] = room
	go room.run()

	w.Write([]byte("ok"))
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")

	mu.Lock()
	room := rooms[roomId]
	mu.Unlock()

	if room == nil {
		http.Error(w, "room not found", 404)
		return
	}

	conn, _ := upgrader.Upgrade(w, r, nil)

	client := &Client{
		conn: conn,
		send: make(chan []byte),
	}

	room.clients[client] = true

	go client.read(room)
	go client.write()
}

func cleanupRooms() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for id, room := range rooms {
			if time.Since(room.createdAt) > 10*time.Minute {
				delete(rooms, id)
			}
		}
		mu.Unlock()
	}
}