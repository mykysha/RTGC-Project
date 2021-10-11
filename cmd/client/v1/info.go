package v1

import (
	"fmt"
	"strings"
)

// GetInfo receives user information.
func (c *Client) GetInfo() error {
	_, err := c.Writer.WriteString("Please enter your 'ID'\n")
	if err != nil {
		return fmt.Errorf("write to CL: %w", err)
	}

	if err = c.Writer.Flush(); err != nil {
		return fmt.Errorf("write to CL: %w", err)
	}

	id, err := c.Reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("read from CL: %w", err)
	}

	c.id = strings.ReplaceAll(id, "\n", "")

	return nil
}
