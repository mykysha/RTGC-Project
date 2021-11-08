package server_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// Warning! Make sure that there are no entities in the database that are used in the test.
// Entities used in the test:
// id 		-	testID1
// username	-	testName1
// roomname	-	testRoom1

type Client struct {
	addr string
	conn *websocket.Conn
}

func start() (*Client, error) {
	c := Client{
		addr: "ws://localhost:8080/app/ws",
		conn: nil,
	}

	conn, _, err := websocket.DefaultDialer.Dial(c.addr, nil)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	c.conn = conn

	return &c, nil
}

func (c Client) send(action string) error {
	r := Request{
		ID:       "testID1",
		Action:   action,
		UserName: "testName1",
		RoomName: "testRoom1",
		Text:     "testText1",
	}

	req, err := encoder(r)
	if err != nil {
		return fmt.Errorf("encoding: %w", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, req)
	if err != nil {
		return fmt.Errorf("writing: %w", err)
	}

	return nil
}

func (c Client) read(readErr chan struct{}) {
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			close(readErr)

			return
		}

		resp, err := decoder(msg)
		if err != nil {
			close(readErr)

			return
		}

		if resp.ID != "testID1" {
			close(readErr)

			return
		}

		if resp.IsError {
			close(readErr)

			return
		}
	}
}

func TestServer(t *testing.T) {
	t.Parallel()

	c, err := start()
	if err != nil {
		t.Fatalf("start: %v", err)
	}

	readErr := make(chan struct{})
	go c.read(readErr)

	time.Sleep(1 * time.Second)

	err = c.send("register")
	if err != nil {
		t.Fatalf("send register: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = c.send("join")
	if err != nil {
		t.Fatalf("send join: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = c.send("send")
	if err != nil {
		t.Fatalf("send register: %v", err)
	}

	time.Sleep(1 * time.Second)

	err = c.send("leave")
	if err != nil {
		t.Fatalf("send register: %v", err)
	}

	time.Sleep(1 * time.Second)

	select {
	case <-readErr:
		t.Fatalf("read err")
	default:
	}
}

type Request struct {
	ID       string `json:"id"`
	Action   string `json:"action"`
	UserName string `json:"uname,omitempty"`
	RoomName string `json:"roomName"`
	Text     string `json:"text,omitempty"`
}

type Response struct {
	IsError     bool   `json:"err"`
	IsMessage   bool   `json:"isMessage"`
	ID          string `json:"id"`
	ErrText     string `json:"errtext,omitempty"`
	MessageText string `json:"messageText,omitempty"`
	FromUser    string `json:"fromUserId,omitempty"`
	FromRoom    string `json:"fromRoom,omitempty"`
}

// encoder to JSON.
func encoder(r Request) ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		err = fmt.Errorf("encode: %w", err)

		return nil, err
	}

	return data, nil
}

// decoder from JSON.
func decoder(msg []byte) (*Response, error) {
	var resp Response

	err := json.Unmarshal(msg, &resp)
	if err != nil {
		err = fmt.Errorf("decode: %w", err)

		return nil, err
	}

	return &resp, nil
}
