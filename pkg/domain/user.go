package domain

type UserCreateRequest struct {
	ID    string   `json:"id"`
	Chats []string `json:"chats"`
}
