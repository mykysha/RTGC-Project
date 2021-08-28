package httpserver

import (
	"net/http"
	"time"
)

type Server struct {
	Httpserver http.Server
}

func MainServer() (mux *http.ServeMux) {
	mux = http.NewServeMux()
	srv := &Server{
		Httpserver: newConfigServer(),
	}
	srv.Httpserver.Handler = http.TimeoutHandler(mux, 100*time.Second, "Timeout!\n")
	return mux
}
