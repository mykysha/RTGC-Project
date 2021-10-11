package app

import (
	"log"

	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
)

type Router struct {
	roomList map[string]*dom.Room
}

func (r *Router) Init() {
	r.roomList = make(map[string]*dom.Room)
}

// NewRoom creates new room.
func (r *Router) NewRoom(userName, roomName string) {
	nr := dom.Room{Name: roomName, UserList: make(map[string]string)}
	r.roomList[roomName] = &nr
	nr.UserList["SERVER"] = "SERVER"

	log.Printf("user '%s' created new room '%s'", userName, roomName)
}

func (r Router) RoomExists(roomName string) bool {
	if _, ok := r.roomList[roomName]; ok {
		return true
	}

	return false
}
