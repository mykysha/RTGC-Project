package domain

import (
	"errors"
	"fmt"
)

// static errors.
var (
	errNoUser = errors.New("no user with such id found in a room")
)

func (r Room) userNameInRoom(userName string) bool {
	if _, ok := r.UserList[userName]; ok {
		return true
	}

	return false
}

func (r Room) idToUserName(userID string) (string, error) {
	var userName string

	for currentName, currentID := range r.UserList {
		if currentID == userID {
			userName = currentName

			return userName, nil
		}
	}

	return "", fmt.Errorf("%w : '%s', '%s'", errNoUser, userID, r.Name)
}
