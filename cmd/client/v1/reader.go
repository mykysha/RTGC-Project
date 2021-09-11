package v1

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Reader gets responses from an open ws-connection.
func Reader(id string, conn *websocket.Conn, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Printf("Read err: %v", err)
			done <- true

			continue
		}

		resp, err := decoder(msg)
		if err != nil {
			log.Printf("Decode err: %v", err)
			done <- true

			continue
		}

		if resp.ID != id {
			log.Printf("ID missmatch")
		}

		if resp.Error {
			log.Printf("Error: %v", resp.Error)
		} else {
			log.Printf("Completed with no errors")
		}

		done <- true
	}
}
