package v1

import (
	"log"

	"github.com/gorilla/websocket"
)

// Dials to websocket connection on server.
func Dial(addr string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	log.Printf("\n"+"Connected to %s", addr)

	return conn
}
