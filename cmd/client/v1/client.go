package v1

import (
	"bufio"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Addr   string
	id     string
	Log    *log.Logger
	conn   *websocket.Conn
	Writer *bufio.Writer
	Reader *bufio.Reader
}

func (c *Client) Init() {
	err := c.GetInfo()
	if err != nil {
		log.Fatalf("userinfo error: %v", err)
	}

	err = c.Dialer()
	if err != nil {
		c.Log.Fatalln(err)
	}

	defer func() {
		err = c.conn.Close()
		if err != nil {
			log.Fatalf("closure error: %v", err)
		}
	}()

	wg := new(sync.WaitGroup)

	activeGoRoutines := 2
	wg.Add(activeGoRoutines)

	go c.infoReader(wg)

	go c.Communicator(wg)

	wg.Wait()
}
