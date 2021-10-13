package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nndergunov/RTGC-Project/api"
)

type Server struct {
	Httpserver http.Server
}

func mainServer() *http.ServeMux {
	var (
		mux     = http.NewServeMux()
		timeout = 10 * time.Second
	)

	srv := &Server{
		Httpserver: *newConfigServer(),
	}
	srv.Httpserver.Handler = http.TimeoutHandler(mux, timeout, "Timeout!\n")

	return mux
}

func New() http.Handler {
	mux := mainServer()
	logger := log.New(os.Stdout, "server ", log.LstdFlags)
	a := &api.API{}

	a.Init(mux, logger)

	return a
}
