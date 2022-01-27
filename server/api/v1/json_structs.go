package v1

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

type State struct {
	State string `json:"state"`
}
