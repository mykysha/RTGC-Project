package server

type Session interface {
	Connection()
	Reader()
	Writer()
	Sender()
}
