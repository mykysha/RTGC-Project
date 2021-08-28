package v1

type UserCreateRequest struct {
	ID    string   `json:"id"`
	Chats []string `json:"chats"`
}
