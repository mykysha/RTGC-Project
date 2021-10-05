package domain

import (
	"errors"
	"fmt"
	"log"
)

var (
	errID    = errors.New("user with such ID already is in the room using valid username")
	errUname = errors.New("such username is already taken in this room")
)

type Room struct {
	Name     string
	UserList map[string]string // username - id
}

// Connecter adds the user to the desired room.
func (r *Room) Connecter(id, userName string) error {
	userNameInRoom, err := r.IDToUserName(id)
	if err == nil {
		return fmt.Errorf("%w : '%v', '%v', '%v'", errID, id, r.Name, userNameInRoom)
	}

	if r.UserNameInRoom(userName) {
		return fmt.Errorf("%w : '%v', '%v'", errUname, userName, r.Name)
	}

	r.UserList[userName] = id

	log.Printf("\n"+"user '%s' connected to the room '%s'", userName, r.Name)

	return nil
}

// Leaver deletes user from the desired room.
func (r *Room) Leaver(userID string) (string, error) {
	userName, err := r.IDToUserName(userID)
	if err != nil {
		return "", err
	}

	delete(r.UserList, userName)
	log.Printf("\n"+"user '%s' disconnected from the room '%s'", userName, r.Name)

	return userName, nil
}

// Messenger gives server list of users in a room that have to receive given message.
func (r Room) Messenger(userID, roomName, text string) (string, string, string, []string, error) {
	returnList := make([]string, 0) // possible error

	userName, err := r.IDToUserName(userID)
	if err != nil {
		return "", "", "", nil, err
	}

	for _, currentID := range r.UserList {
		if currentID == "SERVER" {
			continue
		}

		returnList = append(returnList, currentID)
	}

	return userName, roomName, text, returnList, nil
}
