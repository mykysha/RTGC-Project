package v1

import (
	"encoding/json"
)

type Request struct {
	ID       string `json:"id"`
	Action   string `json:"action"`
	UserName string `json:"uname,omitempty"`
	RoomName string `json:"room_name"`
	Text     string `json:"text,omitempty"`
}

type Response struct {
	ID string `json:"id"`
	Error   bool   `json:"err"`
	ErrText string `json:"errtext,omitempty"`
	IsMessage   bool   `json:"is_message"`
	MessageText string `json:"message_text,omitempty"`
	FromUser    string `json:"from_user_id,omitempty"`
	FromRoom    string `json:"from_room,omitempty"`
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
