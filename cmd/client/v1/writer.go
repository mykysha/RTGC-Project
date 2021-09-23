package v1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

// Writes to an open ws-connection.
func Writer(id string, conn *websocket.Conn, wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()

	fmt.Println("Example command\t 'join:USERNAME:ROOMNAME'")

	for i := 0; ; i++ {
		reader := bufio.NewReader(os.Stdin)

		fmt.Printf("Write command here ->\t")

		msg, _ := reader.ReadString('\n')

		msg = strings.ReplaceAll(msg, "\n", "")

		if !strings.Contains(msg, ":") {
			log.Println("unknown command type")

			continue
		}

		m := strings.Split(msg, ":")
		if len(m) != 3 {
			log.Println("unknown command type")

			continue
		}

		log.Printf("Sending: action - '%s', username - '%s', roomname - '%s'", m[0], m[1], m[2])

		r := Request{ID: id, Action: m[0], Username: m[1], RoomName: m[2]}

		req, err := encode(r)
		if err != nil {
			log.Println("Converting: ", err)

			continue
		}

		err = conn.WriteMessage(websocket.TextMessage, req)
		if err != nil {
			log.Println("Writing: ", err)

			continue
		}

		<-done
	}
}
