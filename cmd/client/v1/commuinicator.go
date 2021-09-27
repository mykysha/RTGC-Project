package v1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// static errors.
var (
	errContain = errors.New("unknown command: doesn`t contain ':'")
	errSplit   = errors.New("unknown command: wrong number of arguments")
	errCom     = errors.New("unknown command")
)

// Communicator handles user-to-server communication.
func Communicator(id string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("\n" + "Possible commands:" +
		"\n\t" + "'join:ROOMNAME:USERNAME'" +
		"\n\t\t" + "(if no room with such name exists, it will be created)" +
		"\n\n\t" + "'send:ROOMNAME:TEXT'" +
		"\n\n\t" + "'leave:ROOMNAME:TEXT'" +
		"\n\t\t" + "(if you don't want to write reason why you leave just type '-')" +
		"\n\t\t" + "(possible reasons: spam\ttoxic community\ttoo many ads\tetc.)")

	for {
		req, err := CL()
		if err != nil {
			log.Printf("\n"+"Request error: %v"+"\n", err)

			continue
		}

		err = WsWriter(id, conn, req)
		if err != nil {
			log.Printf("\n"+"Writing error: %v"+"\n", err)

			continue
		}
	}
}

// CL gets commands from user via Command Line.
func CL() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\n" + "Write command here ->\t")

	msg, _ := reader.ReadString('\n')

	msg = strings.ReplaceAll(msg, "\n", "")

	if !strings.Contains(msg, ":") {
		return nil, errContain
	}

	m := strings.Split(msg, ":")
	if possibleCommandArguments := 3; len(m) != possibleCommandArguments {
		return nil, errSplit
	}

	log.Printf("Sending: action - '%s', '%s', '%s'", m[0], m[1], m[2])

	return m, nil
}

// WsWriter sends requests to an open ws-connection.
func WsWriter(id string, conn *websocket.Conn, m []string) error {
	r := Request{ID: id, Action: m[0], RoomName: m[1]}

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
		ConvErr := fmt.Errorf("converting: %w", err)

		return ConvErr
	}

	err = conn.WriteMessage(websocket.TextMessage, req)
	if err != nil {
		WrErr := fmt.Errorf("writing: %w", err)

		return WrErr
	}

	return nil
}
