package v1

type Request struct {
	ID       string `json:"id"`
	Action   string `json:"action"`
	UserName string `json:"uname,omitempty"`
	RoomName string `json:"room_name"`
	Text     string `json:"text,omitempty"`
}

type Response struct {
	ID          string `json:"id"`
	IsError     bool   `json:"err"`
	ErrText     string `json:"errtext,omitempty"`
	IsMessage   bool   `json:"is_message"`
	MessageText string `json:"message_text,omitempty"`
	FromUser    string `json:"from_user_id,omitempty"`
	FromRoom    string `json:"from_room,omitempty"`
	Time        string `json:"send_time"`
}

type State struct {
	State string `json:"state"`
}
