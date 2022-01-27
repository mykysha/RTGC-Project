package app

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// dialer dials to websocket connection on server.
func (c *Client) dialer() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.addr, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	c.conn = conn

	c.log.Printf("connected to '%s'", c.addr)

	registration := []string{"register", "", ""}

	err = c.wsWriter(registration)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	return nil
}
