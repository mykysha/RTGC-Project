package domain

type Room struct {
	Name     string
	UserList map[string]string // username - id
}
