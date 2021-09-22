package app

import (
	"fmt"
	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
)

func UserNameInRoom(r *dom.Room, userName string) bool {
	if _, ok := r.UserList[userName]; ok {
		return true
	}

	return false
}

func RoomExists(roomName string) bool {
	if _, ok := roomList[roomName]; ok {
		return true
	}

	return false
}

func IDToUserName(r *dom.Room, userID, roomName string) (string, error) {
	var (
		found    bool
		userName string
	)

	for currentName, currentID := range r.UserList {
		if currentID == userID {
			found = true
			userName = currentName

			break
		}
	}

	if !found {
		errNoUser := fmt.Errorf("no user with id '%s' found in room '%s'", userID, roomName)

		return "", errNoUser
	}

	return userName, nil
}
