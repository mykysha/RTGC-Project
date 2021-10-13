package app

import (
	"fmt"
	"sync"
)

// Reader gets responses from an open ws-connection.
func (c Client) infoReader(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			c.log.Printf("reader: %v", err)
		}

		resp, err := decoder(msg)
		if err != nil {
			c.log.Printf("reader: %v", err)
		}

		if resp.ID != c.id {
			c.log.Printf("ID mismatch")
		}

		if resp.IsError {
			c.log.Printf("Error: %v", resp.ErrText)
		} else {
			c.log.Printf("Done")
		}

		if resp.IsMessage {
			message := fmt.Sprintf("\n"+"%s : %s - %s"+"\n", resp.FromRoom, resp.FromUser, resp.MessageText)

			_, err = c.writer.WriteString(message)
			if err != nil {
				c.log.Printf("write to readCommand: %v", err)
			}

			if err = c.writer.Flush(); err != nil {
				c.log.Printf("write to readCommand: %v", err)
			}
		}
	}
}
