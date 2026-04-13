package main

import "time"

type Room struct {
	id        string
	clients   map[*Client]bool
	broadcast chan []byte
	createdAt time.Time
}

func (r *Room) run() {
	for {
		msg := <-r.broadcast
		for c := range r.clients {
			c.send <- msg
		}
	}
}