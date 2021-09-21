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

	if _, ok := r.UserList[userName]; ok {
		errUname := fmt.Errorf("username '%s' already exists in this room", userName)

		return errUname
	}

	r.UserList[userName] = id

	log.Printf("\n"+"User %s connected to the room %s", userName, roomName)

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
	if _, ok := roomList[roomName]; !ok {
		errNoRoom := fmt.Errorf("found no room named %s", roomName)

		return errNoRoom
	}

	var (
		found    bool
		userName string
	)

	r := roomList[roomName]

	for currentName, currentID := range r.UserList {
		if currentID == userID {
			found = true
			userName = currentName

			break
		}
	}

	if !found {
		errNoUser := fmt.Errorf("no user with id %s found in room %s", userID, roomName)

		return errNoUser
	}

	delete(r.UserList, userName)

	return nil
}
