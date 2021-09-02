package main

import (
	"log"
	"net/http"

	"github.com/nndergunov/RTGC-Project/cmd/server/config"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", config.New()))
}
