package chat

type Session interface {
	Connection()
	Reader()
	Writer()
	Sender()
}
