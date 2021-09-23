package app

import (
	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
	"log"
)

// roomList contains Room names and entities.
var roomList = make(map[string]*dom.Room)

// NewRoom creates new room.
func NewRoom(userName, roomName string) {
	nr := dom.Room{Name: roomName, UserList: make(map[string]string)}
	roomList[roomName] = &nr
	nr.UserList["SERVER"] = "SERVER"

	log.Printf("\n"+"user '%s' created new room '%s'", userName, roomName)
}

func RoomExists(roomName string) bool {
	if _, ok := roomList[roomName]; ok {
		return true
	}

	return false
}
