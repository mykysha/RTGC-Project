package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	v1 "github.com/nndergunov/RTGC-Project/api/v1"
	"github.com/nndergunov/RTGC-Project/pkg/app"
)

// API init.

type API struct {
	mux           *http.ServeMux
	log           *log.Logger
	sessions      session
	requestRouter *app.Router
}

type session struct {
	sessionStatus map[*websocket.Conn]bool
	idToSession   map[string]*websocket.Conn
}

func (a *API) Init(m *http.ServeMux, l *log.Logger) {
	a.mux = m
	a.log = l

	a.sessions = session{
		sessionStatus: make(map[*websocket.Conn]bool),
		idToSession:   make(map[string]*websocket.Conn),
	}

	a.requestRouter = &app.Router{}

	a.requestRouter.Init()

	a.mux.HandleFunc("/app/status", a.statusHandler)
	a.mux.HandleFunc("/app/ws", a.wsHandler)
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// statusHandles handles /status request.
func (a API) statusHandler(w http.ResponseWriter, _ *http.Request) {
	resp := v1.State{
		State: "up",
	}

	data, err := statusEncoder(resp)
	if err != nil {
		a.log.Printf("State encoder: %v", err)

		return
	}

	_, err = io.WriteString(w, string(data))
	if err != nil {
		a.log.Printf("State write: %v", err)

		return
	}

	a.log.Printf("Gave status %s", resp.State)
}

// Handles ws connection.
func (a API) wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout:  0,
		ReadBufferSize:    0,
		WriteBufferSize:   0,
		WriteBufferPool:   nil,
		Subprotocols:      nil,
		Error:             nil,
		CheckOrigin:       func(r *http.Request) bool { return true },
		EnableCompression: false,
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.log.Printf("ws fail: %v", err)
	}

	a.sessions.sessionStatus[ws] = true

	defer func(ws *websocket.Conn) {
		delete(a.sessions.sessionStatus, ws)

		err := ws.Close()
		if err != nil {
			a.log.Printf("close fail: %v", err)
		}
	}(ws)

	a.log.Printf("New client connected")

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
			a.log.Printf("reader: %v", err)
			a.errorHandler(ws, "err", true, err)

			continue
		}

		r, err := decode(msg)
		if err != nil {
			a.log.Printf("decoder: %v", err)
			a.errorHandler(ws, "err", true, err)

			continue
		}

		if _, ok := a.sessions.idToSession[r.ID]; !ok {
			a.sessions.idToSession[r.ID] = ws
		}

		err = a.communicator(r)
		if err != nil {
			a.errorHandler(ws, r.ID, true, err)
		}
	}
}

func (a API) communicator(r v1.Request) error {
	fromUser, fromRoom, message, toID, err := a.requestRouter.ActionHandler(r.ID, r.Action, r.RoomName, r.UserName, r.Text)
	if err != nil {
		return fmt.Errorf("communicator: %w", err)
	}

	wgSender := new(sync.WaitGroup)

	for _, id := range toID {
		wgSender.Add(1)

		go a.sender(a.sessions.idToSession[id], id, fromUser, fromRoom, message)

		wgSender.Done()
	}

	wgSender.Wait()

	return nil
}

// responser sends completion status to the client.
func (a API) errorHandler(ws *websocket.Conn, id string, e bool, err error) {
	resp := v1.Response{
		IsError:     e,
		IsMessage:   false,
		ID:          id,
		ErrText:     fmt.Sprintf("%v", err),
		MessageText: "",
		FromUser:    "",
		FromRoom:    "",
	}

	msg, err := encode(resp)
	if err != nil {
		a.log.Printf("errorHandler: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.log.Printf("errorHandler: %v", err)

		return
	}
}

// sender sends message to desired user.
func (a API) sender(ws *websocket.Conn, id, fromUser, fromRoom, message string) {
	resp := v1.Response{
		IsError:     false,
		IsMessage:   true,
		ID:          id,
		ErrText:     "",
		MessageText: message,
		FromUser:    fromUser,
		FromRoom:    fromRoom,
	}

	msg, err := encode(resp)
	if err != nil {
		a.log.Printf("errorWriter: %v", err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.log.Print(err)
	}
}
