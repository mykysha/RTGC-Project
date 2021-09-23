package v1

type Request struct {
	ID       string `json:"id"`
	Action   string `json:"action"`
	Username string `json:"uname"`
	RoomName string `json:"chatname"`
}
