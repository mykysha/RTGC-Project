package app

import (
	"fmt"
	"strings"
)

// getInfo receives user information.
func (c *Client) getInfo() error {
	_, err := c.writer.WriteString("Please enter your 'ID'\n")
	if err != nil {
		return fmt.Errorf("write to readCommand: %w", err)
	}

	if err = c.writer.Flush(); err != nil {
		return fmt.Errorf("write to readCommand: %w", err)
	}

	id, err := c.reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("read from readCommand: %w", err)
	}

	c.id = strings.ReplaceAll(id, "\n", "")

	return nil
}
