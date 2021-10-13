package app

import (
	"bufio"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	addr   string
	id     string
	log    *log.Logger
	conn   *websocket.Conn
	writer *bufio.Writer
	reader *bufio.Reader
}

func (c *Client) Init(a string, l *log.Logger, w *bufio.Writer, r *bufio.Reader) {
	c.addr = a
	c.log = l
	c.writer = w
	c.reader = r

	err := c.getInfo()
	if err != nil {
		log.Fatalf("userinfo error: %v", err)
	}

	err = c.dialer()
	if err != nil {
		c.log.Fatalln(err)
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

	go c.communicator(wg)

	wg.Wait()
}
