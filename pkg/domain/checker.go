package domain

import (
	"fmt"
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
		errNoUser := fmt.Errorf("no user with id '%s' found in room '%s'", userID, r.Name)

		return "", errNoUser
	}

	return userName, nil
}
