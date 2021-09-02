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
	srv := &Server{
		Httpserver: *newConfigServer(),
	}
	srv.Httpserver.Handler = http.TimeoutHandler(mux, 100*time.Second, "Timeout!\n")

	return mux
}

func New() http.Handler {
	mux := MainServer()
	log := log.New(os.Stdout, "server ", log.LstdFlags)
	a := &api.API{
		Mux: mux,
		Log: log,
	}
	a.Init()

	return a
}
