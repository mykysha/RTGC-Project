package main

import (
	"bufio"
	"log"
	"os"

	v1 "github.com/nndergunov/RTGC-Project/cmd/client/app"
)

const addr = "ws://localhost:8080/app/ws"

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)

	l := log.New(os.Stdout, "client ", log.LstdFlags)

	c := v1.Client{}

	c.Init(addr, l, w, r)
}
