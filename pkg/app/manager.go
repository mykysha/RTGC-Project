package app

import (
	"fmt"
	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
	"log"
)

// roomList contains Room names and entities.
var roomList = make(map[string]*dom.Room)

// Connecter adds the user to the desired room.
func Connecter(id, userName, roomName string) error {
	if _, ok := roomList[roomName]; !ok {
		NewRoom(userName, roomName)
	}

	r := roomList[roomName]

	if UserNameInRoom(r, userName) {
		errUname := fmt.Errorf("username '%s' already exists in this room", userName)

		return errUname
	}

	r.UserList[userName] = id

	log.Printf("\n"+"user %s connected to the room %s", userName, roomName)

	return nil
}

// NewRoom creates new room.
func NewRoom(userName, roomName string) {
	nr := dom.Room{Name: roomName, UserList: make(map[string]string)}
	roomList[roomName] = &nr

	log.Printf("\n"+"user %s created new room %s", userName, roomName)
}

// Leaver deletes user from the desired room.
func Leaver(userID, roomName string) error {
	if !RoomExists(roomName) {
		errNoRoom := fmt.Errorf("found no room named %s", roomName)

		return errNoRoom
	}

	r := roomList[roomName]

	userName, findErr := IDToUserName(r, userID, roomName)
	if findErr != nil {
		return findErr
	}

	delete(r.UserList, userName)

	return nil
}
