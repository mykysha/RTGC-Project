package main
 
import (
    "log"
 
    "github.com/gorilla/websocket"
)

const addr = "ws://localhost:8080/v1/ws"

func main() {
    conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
    if err != nil {
        log.Fatal("connection error:", err)
    }
    log.Printf("connected to %s", addr)
    defer conn.Close()
}