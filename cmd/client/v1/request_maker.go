package v1

import (
	"encoding/json"
	"log"
)

type Request struct {
	ID       string `json:"id"`
	Action   string `json:"action"`
	Username string `json:"uname"`
	RoomName string `json:"chatname"`
}

type Response struct {
	ID      string `json:"id"`
	Error   bool   `json:"err"`
	ErrText string `json:"errtext,omitempty"`
}

func encode(r Request) ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return data, nil
}

func decode(msg []byte) (*Response, error) {
	var resp Response

	err := json.Unmarshal(msg, &resp)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &resp, nil
}
