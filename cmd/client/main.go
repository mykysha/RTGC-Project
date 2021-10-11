package main

import (
	"bufio"
	"log"
	"os"

	v1 "github.com/nndergunov/RTGC-Project/cmd/client/v1"
)

const addr = "ws://localhost:8080/v1/ws"

func main() {
	writer := bufio.NewWriter(os.Stdout)
	reader := bufio.NewReader(os.Stdin)

	logger := log.New(os.Stdout, "client ", log.LstdFlags)

	c := v1.Client{
		Addr:   addr,
		Log:    logger,
		Writer: writer,
		Reader: reader,
	}

	c.Init()
}
