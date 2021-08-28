package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (a *API) wsHandler(w http.ResponseWriter, r *http.Request) {

        ws, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                log.Printf("ws fail: %v", err)
        }
		log.Println("ws success")
		ws.WriteMessage(websocket.TextMessage, []byte("Init correct"))
		defer ws.Close()
}
