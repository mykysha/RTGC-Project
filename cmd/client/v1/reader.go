package v1

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
			c.Log.Printf("reader: %v", err)
		}

		resp, err := decoder(msg)
		if err != nil {
			c.Log.Printf("reader: %v", err)
		}

		if resp.ID != c.id {
			c.Log.Printf("ID mismatch")
		}

		if resp.IsError {
			c.Log.Printf("Error: %v", resp.ErrText)
		} else {
			c.Log.Printf("Done")
		}

		if resp.IsMessage {
			message := fmt.Sprintf("\n"+"%s : %s - %s"+"\n", resp.FromRoom, resp.FromUser, resp.MessageText)

			_, err = c.Writer.WriteString(message)
			if err != nil {
				c.Log.Printf("write to CL: %v", err)
			}

			if err = c.Writer.Flush(); err != nil {
				c.Log.Printf("write to CL: %v", err)
			}
		}
	}
}
