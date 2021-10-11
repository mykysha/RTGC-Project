package v1

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Dialer dials to websocket connection on server.
func (c *Client) Dialer() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.Addr, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	c.conn = conn

	c.Log.Printf("connected to '%s'", c.Addr)

	return nil
}
