package chat

type Chat interface {
	New()
	Join()
	Send()
	Leave()
	Users()
	Viewer()
}