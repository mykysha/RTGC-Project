package v1

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// static errors.
var (
	errContain = errors.New("unknown command: does not contain ':'")
	errSplit   = errors.New("unknown command: wrong number of arguments")
	errCom     = errors.New("unknown command")
)

// Communicator handles user-to-server communication.
func (c Client) Communicator(wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := c.Writer.WriteString("\n" + "Possible commands:" +
		"\n\t" + "'join:ROOMNAME:USERNAME'" +
		"\n\t\t" + "(if no room with such name exists, it will be created)" +
		"\n\n\t" + "'send:ROOMNAME:TEXT'" +
		"\n\n\t" + "'leave:ROOMNAME:TEXT'" +
		"\n\t\t" + "(if you don't want to write reason why you leave just type '-')" +
		"\n\t\t" + "(possible reasons: spam\ttoxic community\ttoo many ads\tetc.)" + "\n")
	if err != nil {
		c.Log.Printf("write to CL: %v", err)
	}

	if err = c.Writer.Flush(); err != nil {
		c.Log.Printf("write to CL: %v", err)
	}

	for {
		req, err := c.CL()
		if err != nil {
			c.Log.Printf("Request error: %v", err)

			continue
		}

		err = c.WsWriter(req)
		if err != nil {
			c.Log.Printf("Writing error: %v", err)

			continue
		}
	}
}

// CL gets commands from user via Command Line.
func (c Client) CL() ([]string, error) {
	_, err := c.Writer.WriteString("\n" + "Write command in command line:" + "\n")
	if err != nil {
		return nil, fmt.Errorf("write to CL: %w", err)
	}

	if err = c.Writer.Flush(); err != nil {
		return nil, fmt.Errorf("write to CL: %w", err)
	}

	msg, err := c.Reader.ReadString('\n')
	if err != nil {
		return nil, fmt.Errorf("read from CL: %w", err)
	}

	msg = strings.ReplaceAll(msg, "\n", "")

	if !strings.Contains(msg, ":") {
		return nil, errContain
	}

	m := strings.Split(msg, ":")
	if possibleCommandArguments := 3; len(m) != possibleCommandArguments {
		return nil, errSplit
	}

	c.Log.Printf("Sending: action - '%s', '%s', '%s'", m[0], m[1], m[2])

	return m, nil
}

// WsWriter sends requests to an open ws-connection.
func (c Client) WsWriter(m []string) error {
	r := Request{
		ID:       c.id,
		Action:   m[0],
		RoomName: m[1],
		UserName: "",
		Text:     "",
	}

	switch r.Action {
	case "join":
		r.UserName = m[2]

	case "send", "leave":
		r.Text = m[2]

	default:
		return errCom
	}

	req, err := encoder(r)
	if err != nil {
		err = fmt.Errorf("converting: %w", err)

		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, req)
	if err != nil {
		err = fmt.Errorf("writing: %w", err)

		return err
	}

	return nil
}
