package main

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) read(room *Room) {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}
		room.broadcast <- msg
	}
}

func (c *Client) write() {
	for msg := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, msg)
	}
}