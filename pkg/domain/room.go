package domain

import (
	"fmt"
	"log"
)

type Room struct {
	Name     string
	UserList map[string]string // username - id
}

// Connecter adds the user to the desired room.
func (r *Room) Connecter(id, userName string) error {
	userNameInRoom, isInErr := r.IDToUserName(id)
	if isInErr == nil {
		errID := fmt.Errorf("user with id '%s' "+
			"is already connected to the room '%s' "+
			"under the username '%s'", id, r.Name, userNameInRoom)

		return errID
	}

	if r.UserNameInRoom(userName) {
		errUname := fmt.Errorf("username '%s' already exists in this room", userName)

		return errUname
	}

	r.UserList[userName] = id

	log.Printf("\n"+"user '%s' connected to the room '%s'", userName, r.Name)

	return nil
}

// Leaver deletes user from the desired room.
func (r *Room) Leaver(userID string) (string, error) {
	userName, findErr := r.IDToUserName(userID)
	if findErr != nil {
		return "", findErr
	}

	delete(r.UserList, userName)
	log.Printf("\n"+"user '%s' disconnected from the room '%s'", userName, r.Name)

	return userName, nil
}

// Messenger gives server list of users in a room that have to receive given message.
func (r Room) Messenger(userID, roomName, text string) (string, string, string, []string, error) {
	var returnList []string

	userName, findErr := r.IDToUserName(userID)

	if findErr != nil {
		return "", "", "", nil, findErr
	}

	for _, currentID := range r.UserList {
		if currentID == "SERVER" {
			continue
		}

		returnList = append(returnList, currentID)
	}

	return userName, roomName, text, returnList, nil
}
