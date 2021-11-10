package config

import (
	"net/http"
	"time"
)

type Server struct {
	Httpserver http.Server
}

func MainServer() *http.ServeMux {
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
