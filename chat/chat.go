package chat

type Chat interface {
	Join()
	Leave()
	Users()
	Viewer()
}