package v1

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Reader gets responses from an open ws-connection.
func Reader(id string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			log.Printf("\n"+"Read err: %v"+"\n", err)

			continue
		}

		resp, err := decoder(msg)
		if err != nil {
			log.Printf("\n"+"Decode err: %v"+"\n", err)

			continue
		}

		if resp.ID != id {
			log.Printf("\n" + "ID missmatch" + "\n")
		}

		if resp.Error {
			log.Printf("Error: %v", resp.ErrText)
		} else {
			log.Printf("\n" + "Completed with no errors" + "\n")
		}

		if resp.IsMessage {
			fmt.Printf("\n"+"%s : %s - %s"+"\n", resp.FromRoom, resp.FromUser, resp.MessageText)
		}
	}
}
