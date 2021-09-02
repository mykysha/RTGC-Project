package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// API init.

type API struct {
	Mux *http.ServeMux
	Log *log.Logger
}

func (a API) Init() {
	a.Mux.HandleFunc("/v1/status", a.statusHandler)
	a.Mux.HandleFunc("/v1/ws", a.wsHandler)
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Mux.ServeHTTP(w, r)
}

// /status handler.

type Status struct {
	State string `json:"state"`
}

func (a *API) statusHandler(w http.ResponseWriter, r *http.Request) {
	response := Status{
		State: "up",
	}

	err := encode(w, response)
	if err != nil {
		log.Printf("json fail: %v", err)

		return
	}
}

// /ws handler.

var Sessions = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{}

func (a *API) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ws fail: %v", err)
	}

	Sessions[ws] = true

	log.Println("ws connection succesfull")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Init correct"))
	if err != nil {
		log.Printf("ws fail: %v", err)
	}

	defer sessionClose(ws)
	defer ws.Close()
}

func sessionClose(ws *websocket.Conn) {
	Sessions[ws] = false
}
