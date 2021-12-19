package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nndergunov/RTGC-Project/client/app"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	r := bufio.NewReader(os.Stdin)

	l := log.New(os.Stdout, "client ", log.LstdFlags)

	c := app.Client{}

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("env file read: %v", err)
	}

	addr := fmt.Sprintf(
		os.Getenv("PROTOCOL") +
			"://" + os.Getenv("HOST") +
			":" + os.Getenv("PORT") +
			os.Getenv("ENDPOINT"),
	)

	c.Init(addr, l, w, r)
}
