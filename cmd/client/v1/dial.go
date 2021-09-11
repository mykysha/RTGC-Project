package v1

import (
	"log"

	"github.com/gorilla/websocket"
)

// Dialer dials to websocket connection on server.
func Dialer(addr string) *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("connection error:", err)
	}

	log.Printf("\n"+"connected to %s", addr)

	return conn
}
