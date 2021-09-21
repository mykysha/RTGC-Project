package app

import "fmt"

// Messenger gives server list of users in a room that have to receive message.
func Messenger(userID, roomName, text string) (string, string, string, []string, error) {
	if _, ok := roomList[roomName]; !ok {
		errNoRoom := fmt.Errorf("found no room named %s", roomName)

		return "", "", "", nil, errNoRoom
	}

	r := roomList[roomName]

	var (
		found      bool
		userName   string
		returnList []string
	)

	for currentName, currentID := range r.UserList {
		if currentID == userID {
			found = true
			userName = currentName

			break
		}
	}

	if !found {
		errNoUser := fmt.Errorf("no user with id %s found in room %s", userID, roomName)

		return "", "", "", nil, errNoUser
	}

	for _, currentID := range r.UserList {
		returnList = append(returnList, currentID)
	}

	return userName, roomName, text, returnList, nil
}
