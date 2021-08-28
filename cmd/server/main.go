package server

import (
	"net/http"

	"github.com/nndergunov/RTGC-Project/api"
	"github.com/nndergunov/RTGC-Project/internal/httpserver"
)

func Main() {
	http.ListenAndServe(":8080", New())
}

func New() http.Handler {
	mux := httpserver.MainServer()
	a := &api.API{Mux: mux}
	a.Init()
	return a
}