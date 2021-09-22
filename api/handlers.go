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
			a.errorWriter(ws, "err", err)

			continue
		}

		r, err := decode(msg)
		if err != nil {
			a.Log.Printf("decoder: %v", err)
			a.errorWriter(ws, "err", err)

			continue
		}

		if _, ok := IDSessions[r.ID]; !ok {
			IDSessions[r.ID] = ws
		}

		actionErr := a.actionHandler(r)
		if actionErr != nil {
			a.errorWriter(ws, r.ID, actionErr)
		} else {
			a.indicator(ws, r.ID)
		}
	}
}

// errorWriter sends errors to the client.
func (a API) errorWriter(ws *websocket.Conn, id string, err error) {
	resp := v1.Response{ID: id, Error: true, ErrText: fmt.Sprintf("%v", err)}

	msg, err := encode(resp)
	if err != nil {
		a.Log.Printf("errorWriter: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Print(err)

		return
	}
}

func (a API) indicator(ws *websocket.Conn, id string) {
	resp := v1.Response{ID: id, Error: false}

	msg, err := encode(resp)
	if err != nil {
		a.Log.Printf("indicator: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Printf("indicator: %v", err)

		return
	}
}

// sender sends message to desired user.
func (a API) sender(ws *websocket.Conn, id, fromUser, fromRoom, message string) {
	resp := v1.Response{
		ID:          id,
		Error:       false,
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

		return
	}
}

func (a API) actionHandler(r v1.Request) error {
	switch r.Action {
	case "join":
		joinErr := a.joinHandler(r)

		return joinErr

	case "send":
		sendErr := a.sendHandler(r)

		return sendErr

	case "leave":
		leaveErr := a.leaveHandler(r)

		return leaveErr

	default:
		unknownAction := fmt.Errorf("action '%s' not supported", r.Action)

		return unknownAction
	}
}

// joinHandler handles join request.
func (a API) joinHandler(r v1.Request) error {
	log.Printf("\n"+"ID: '%s', Action: '%s', UserName: '%s', RoomName: '%s'", r.ID, r.Action, r.UserName, r.RoomName)
	conErr := app.Connecter(r.ID, r.UserName, r.RoomName)

	return conErr
}

// sendHandler handles send request.
func (a API) sendHandler(r v1.Request) error {
	log.Printf("\n"+"ID: '%s', Action: '%s', RoomName: '%s', Text: '%s'", r.ID, r.Action, r.RoomName, r.Text)

	fromUser, fromRoom, message, toID, messageErr := app.Messenger(r.ID, r.RoomName, r.Text)
	if messageErr != nil {
		return messageErr
	}

	wgSender := new(sync.WaitGroup)

	for _, id := range toID {
		wgSender.Add(1)

		go a.sender(IDSessions[id], id, fromUser, fromRoom, message)
	}

	wgSender.Wait()

	return nil
}

// leaveHandler handles leave request.
func (a API) leaveHandler(r v1.Request) error {
	log.Printf("\n"+"ID: '%s', Action: '%s', RoomName: '%s', Text: '%s'", r.ID, r.Action, r.RoomName, r.Text)

	if r.Text != "-" {
		sendErr := a.sendHandler(r)
		if sendErr != nil {
			return sendErr
		}
	}

	leaveErr := app.Leaver(r.ID, r.RoomName)
	if leaveErr != nil {
		return leaveErr
	}

	return nil
}
