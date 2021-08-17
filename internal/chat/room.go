package chat

type Room interface {
	New()
	Join()
	Broadcast()
	Leave()
	Users()
	Viewer()
}
