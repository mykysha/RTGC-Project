package domain

import (
	"errors"
	"fmt"
)

// static errors.
var (
	errNoUser = errors.New("no user with such id found in a room")
)

func (r Room) UserNameInRoom(userName string) bool {
	if _, ok := r.UserList[userName]; ok {
		return true
	}

	return false
}

func (r Room) IDToUserName(userID string) (string, error) {
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
		return "", fmt.Errorf("%w : '%s', '%s'", errNoUser, userID, r.Name)
	}

	return userName, nil
}
