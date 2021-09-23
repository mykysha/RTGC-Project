package api

import (
	"fmt"
	"github.com/nndergunov/RTGC-Project/pkg/app"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nndergunov/RTGC-Project/api/v1"
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
var IDSessions = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{}

// Handles ws connection.
func (a API) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.Log.Printf("ws fail: %v", err)
	}

	Sessions[ws] = true

	defer func(ws *websocket.Conn) {
		delete(Sessions, ws)

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
			a.responser(ws, "err", true, err)

			continue
		}

		r, err := decode(msg)
		if err != nil {
			a.Log.Printf("decoder: %v", err)
			a.responser(ws, "err", true, err)

			continue
		}

		if _, ok := IDSessions[r.ID]; !ok {
			IDSessions[r.ID] = ws
		}

		actionErr := a.communicator(r)
		if actionErr != nil {
			a.responser(ws, r.ID, true, actionErr)
		} else {
			a.responser(ws, r.ID, false, nil)
		}
	}
}

func (a API) communicator(r v1.Request) error {
	fromUser, fromRoom, message, toID, actionErr := app.ActionHandler(r.ID, r.Action, r.RoomName, r.UserName, r.Text)
	if actionErr != nil {
		return actionErr
	}

	wgSender := new(sync.WaitGroup)

	for _, id := range toID {

		wgSender.Add(1)

		go a.sender(IDSessions[id], id, fromUser, fromRoom, message)

		wgSender.Done()
	}

	wgSender.Wait()

	return nil
}

// responser sends completion status to the client.
func (a API) responser(ws *websocket.Conn, id string, e bool, err error) {
	resp := v1.Response{ID: id}
	if e {
		resp.IsError = true
		resp.ErrText = fmt.Sprintf("%v", err)
	} else {
		resp.IsError = false
	}

	msg, err := encode(resp)
	if err != nil {
		a.Log.Printf("responser: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Printf("responser: %v", err)

		return
	}
}

// sender sends message to desired user.
func (a API) sender(ws *websocket.Conn, id, fromUser, fromRoom, message string) {
	resp := v1.Response{
		ID:          id,
		IsError:     false,
		IsMessage:   true,
		MessageText: message,
		FromUser:    fromUser,
		FromRoom:    fromRoom,
	}

	msg, err := encode(resp)
	if err != nil {
		a.Log.Printf("errorWriter: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Print(err)
	}
}
