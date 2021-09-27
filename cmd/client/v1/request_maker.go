package v1

import (
	"encoding/json"
)

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
		return nil, err
	}

	return data, nil
}

// decoder from JSON.
func decoder(msg []byte) (*Response, error) {
	var resp Response

	err := json.Unmarshal(msg, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
