package app

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// communicator handles user-to-server communication.
func (c Client) communicator(wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := c.writer.WriteString("\n" + "Possible commands:" +
		"\n\t" + "'join:ROOMNAME:USERNAME'" +
		"\n\t\t" + "(if no room with such name exists, it will be created)" +
		"\n\n\t" + "'send:ROOMNAME:TEXT'" +
		"\n\n\t" + "'leave:ROOMNAME:TEXT'" +
		"\n\t\t" + "(if you don't want to write reason why you leave just type '-')" +
		"\n\t\t" + "(possible reasons: spam\ttoxic community\ttoo many ads\tetc.)" + "\n")
	if err != nil {
		c.log.Printf("write to readCommand: %v", err)
	}

	if err = c.writer.Flush(); err != nil {
		c.log.Printf("write to readCommand: %v", err)
	}

	for {
		req, err := c.readCommand()
		if err != nil {
			c.log.Printf("Request error: %v", err)

			continue
		}

		err = c.wsWriter(req)
		if err != nil {
			c.log.Printf("Writing error: %v", err)

			continue
		}
	}
}

// readCommand gets commands from user via Command Line.
func (c Client) readCommand() ([]string, error) {
	_, err := c.writer.WriteString("\n" + "Write command in command line:" + "\n")
	if err != nil {
		return nil, fmt.Errorf("write to readCommand: %w", err)
	}

	if err = c.writer.Flush(); err != nil {
		return nil, fmt.Errorf("write to readCommand: %w", err)
	}

	var msg string

	for {
		msg, err = c.reader.ReadString('\n')
		if err == nil {
			break
		}

		time.Sleep(time.Second)
	}

	msg = strings.ReplaceAll(msg, "\n", "")

	if !strings.Contains(msg, ":") {
		return nil, errContain
	}

	m := strings.Split(msg, ":")
	if possibleCommandArguments := 3; len(m) != possibleCommandArguments {
		return nil, errSplit
	}

	if m[0] == "register" {
		return nil, fmt.Errorf("%w: %s", errUnauthUse, m[0])
	}

	c.log.Printf("Sending: action - '%s', '%s', '%s'", m[0], m[1], m[2])

	return m, nil
}

// wsWriter sends requests to an open ws-connection.
func (c Client) wsWriter(m []string) error {
	r := Request{
		ID:       c.id,
		Action:   m[0],
		RoomName: m[1],
		UserName: "",
		Text:     "",
	}

	switch r.Action {
	case "register":
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