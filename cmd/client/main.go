package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
	client "github.com/nndergunov/RTGC-Project/cmd/client/v1"
)

const addr = "ws://localhost:8080/v1/ws"

func main() {
	id, err := client.GetInfo()
	if err != nil {
		log.Fatalf("userinfo error: %v", err)
	}

	conn := client.Dialer(addr)

	defer func(conn *websocket.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatalf("closure error: %v", err)
		}
	}(conn)

	wg := new(sync.WaitGroup)

	activeGoRoutines := 2
	wg.Add(activeGoRoutines)

	go func() {
		err = client.Reader(id, conn, wg)
		if err != nil {
			log.Println(err)
		}
	}()
	go client.Communicator(id, conn, wg)

	wg.Wait()
}
