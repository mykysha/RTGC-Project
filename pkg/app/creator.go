package app

import (
	"log"

	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
)

// RoomList contains Room names and entities.
var RoomList = make(map[string]*dom.Room)

// Creates new room.
func NewRoom(userName, roomName string) {
	nr := dom.Room{Name: roomName, UserList: make([]string, 1)}
	RoomList[roomName] = &nr

	log.Printf("\n"+"User %s created new room %s", userName, roomName)
}
