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

func MainServer() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	timeout := 100 * time.Second

	srv := &Server{
		Httpserver: *newConfigServer(),
	}
	srv.Httpserver.Handler = http.TimeoutHandler(mux, timeout, "Timeout!\n")

	return mux
}

func New() http.Handler {
	mux := MainServer()
	logger := log.New(os.Stdout, "server ", log.LstdFlags)
	a := &api.API{
		Mux: mux,
		Log: logger,
	}
	a.Init()

	return a
}
