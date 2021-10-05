package v1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

// Reader gets responses from an open ws-connection.
func Reader(id string, conn *websocket.Conn, wg *sync.WaitGroup) error {
	defer wg.Done()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("reader: %w", err)
		}

		resp, err := decoder(msg)
		if err != nil {
			return fmt.Errorf("reader: %w", err)
		}

		if resp.ID != id {
			log.Printf("\n" + "ID missmatch" + "\n")
		}

		if resp.IsError {
			log.Printf("Error: %v", resp.ErrText)
		} else {
			log.Printf("\n" + "Done" + "\n")
		}

		if resp.IsMessage {
			writer := bufio.NewWriter(os.Stdout)
			message := fmt.Sprintf("\n"+"%s : %s - %s"+"\n", resp.FromRoom, resp.FromUser, resp.MessageText)

			_, err = writer.WriteString(message)
			if err != nil {
				return fmt.Errorf("write to CL: %w", err)
			}

			if err = writer.Flush(); err != nil {
				return fmt.Errorf("write to CL: %w", err)
			}
		}
	}
}
