package main

import (
	"sync"

	client "github.com/nndergunov/RTGC-Project/cmd/client/v1"
)

const addr = "ws://localhost:8080/v1/ws"

func main() {
	id := client.GetInfo()

	conn := client.Dial(addr)
	defer conn.Close()

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go client.Reader(id, conn, wg)
	go client.Writer(id, conn, wg)

	wg.Wait()
}
