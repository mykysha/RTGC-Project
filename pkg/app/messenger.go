package app

import "fmt"

// Messenger gives server list of users in a room that have to receive message.
func Messenger(userID, roomName, text string) (string, string, string, []string, error) {
	if !RoomExists(roomName) {
		errNoRoom := fmt.Errorf("found no room named '%s'", roomName)

		return "", "", "", nil, errNoRoom
	}

	r := roomList[roomName]

	var returnList []string

	userName, findErr := IDToUserName(r, userID, roomName)

	if findErr != nil {
		return "", "", "", nil, findErr
	}

	for _, currentID := range r.UserList {
		returnList = append(returnList, currentID)
	}

	return userName, roomName, text, returnList, nil
}
