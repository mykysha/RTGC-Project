package v1

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Reads from an open ws-connection.
func Reader(id string, conn *websocket.Conn, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println("Read err: ", err)
			done <- true

			continue
		}

		resp, err := decode(msg)
		if err != nil {
			log.Println("Decode err: ", err)
			done <- true

			continue
		}

		if resp.Error {
			log.Printf("\n"+"Error: %v", resp.Error)
		} else {
			log.Println("Completed with no errors")
		}
		done <- true
	}
}
