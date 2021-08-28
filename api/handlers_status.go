package api

import "net/http"

type Status struct {
	State string `json:"state"`
}

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	response := Status{
		State: "up",
	}
	encode(w, response)
}
