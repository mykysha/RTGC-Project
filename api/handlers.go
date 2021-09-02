package api

import (
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

// /status handler.


func (a API) statusHandler(w http.ResponseWriter, r *http.Request) {
	response := v1.Status{
		State: "up",
	}

	data, err := statusEncoder(response)
	if err != nil {
		a.Log.Println(err)
	}

	_, err = io.WriteString(w, string(data))
	if err != nil {
		a.Log.Println(err)
	}

	a.Log.Println("Gave status")
}

// /ws handler.

var Sessions = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{}

func (a API) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		a.Log.Println("ws fail:", err)
	}

	Sessions[ws] = true

	defer sessionClose(ws)

	a.Log.Println("New client")

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go a.reader(ws, wg)

	wg.Wait()
}

func sessionClose(ws *websocket.Conn) {
	Sessions[ws] = false

	ws.Close()
}

func (a API) reader(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			a.Log.Println(err)

			return
		}

		r, err := decode(msg)
		if err != nil {
			a.Log.Println(err)

			return
		}

		log.Printf("\n" + "ID: %s, Action: %s, Username: %s, RoomName: %s", r.ID, r.Action, r.Username, r.RoomName)
	}
}

func (a API) writer(ws *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	resp := v1.Response{ID: "your id", Error: false}
	msg, err := encode(resp)
	if err != nil {
		a.Log.Println(err)

		return
	}

	err = ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		a.Log.Println(err)

		return
		}
	}
