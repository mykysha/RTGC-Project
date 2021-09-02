package v1

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Reads from an open ws-connection.
func Reader(id string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read: ", err)

			return
		}

		resp, err := decode(msg)
		if err != nil {
			log.Println("Decode: ", err)

			return
		}

		fmt.Println(resp.Error)

	}
}
