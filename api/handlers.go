package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nndergunov/RTGC-Project/api/v1"
	"github.com/nndergunov/RTGC-Project/pkg/app"
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

// /status.
func (a API) statusHandler(w http.ResponseWriter, _ *http.Request) {
	resp := v1.State{
		State: "up",
	}

	data, err := statusEncoder(resp)
	if err != nil {
		a.Log.Printf("State encoder: %v", err)
		return
	}

	_, err = io.WriteString(w, string(data))
	if err != nil {
		a.Log.Printf("State write: %v", err)
		return
	}

	a.Log.Printf("Gave status %s", resp.State)
}

// /ws.

var Sessions = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{}

// Handles ws connection.
func (a API) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.Log.Printf("ws fail: %v", err)
	}

	Sessions[ws] = true

	defer func(ws *websocket.Conn) {
		Sessions[ws] = false

		err := ws.Close()
		if err != nil {
			a.Log.Printf("close fail: %v", err)
		}
	}(ws)

	a.Log.Printf("New client added")

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go a.reader(ws, wg)

	wg.Wait()
}

// Gets requests from the client.
func (a API) reader(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			a.Log.Printf("reader: %v", err)
			a.writer(ws, "err", true, err)

			continue
		}

		r, err := decode(msg)
		if err != nil {
			a.Log.Printf("decoder: %v", err)
			a.writer(ws, "err", true, err)

			continue
		}

		log.Printf("\n"+"ID: %s, Action: %s, UserName: %s, RoomName: %s", r.ID, r.Action, r.UserName, r.RoomName)

		switch r.Action {
		case "join":
			app.Connecter(r.ID, r.UserName, r.RoomName)
		case "send":
			app.Messenger(r.ID, r.RoomName, r.Text)
		case "leave":
			if r.Text != "-" {
				app.Messenger(r.ID, r.RoomName, r.Text)
			}
			app.Leaver(r.UserName, r.RoomName)
		default:
			a.writer(ws, r.ID, true, fmt.Errorf("action not supported"))
		}

		a.writer(ws, r.ID, false, nil)
	}
}

// Responds to the client.
func (a API) writer(ws *websocket.Conn, id string, e bool, err error) {
	resp := v1.Response{ID: id, Error: e, ErrText: fmt.Sprintf("%v", err)}

	msg, err := encode(resp)
	if err != nil {
		a.Log.Printf("writer: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Print(err)

		return
	}
}
