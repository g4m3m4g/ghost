package client

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func JoinRoom(roomId string) {
	url := "ws://localhost:8080/ws/" + roomId

	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("failed to connect")
		return
	}

	// receive
	go func() {
		for {
			_, msg, _ := conn.ReadMessage()
			fmt.Println(string(msg))
		}
	}()

	// send
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		conn.WriteMessage(websocket.TextMessage, scanner.Bytes())
	}
}