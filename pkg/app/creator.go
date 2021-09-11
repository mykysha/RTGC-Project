package app

import (
	"log"

	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
)

// RoomList contains Room names and entities.
var RoomList = make(map[string]*dom.Room)

// NewRoom creates new room.
func NewRoom(userName, roomName string) {
	// TODO check if roomname is already taken
	nr := dom.Room{Name: roomName, UserList: make(map[string]string)}
	RoomList[roomName] = &nr

	log.Printf("\n"+"User %s created new room %s", userName, roomName)
}
