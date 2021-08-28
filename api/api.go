package api

import (
	"net/http"
)

type API struct {
	Mux *http.ServeMux
}

func (a *API) Init() {
	a.Mux.HandleFunc("/v1/status", a.statusHandler)
	a.Mux.HandleFunc("/v1/ws", a.wsHandler)
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    a.Mux.ServeHTTP(w, r)
}