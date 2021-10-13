package domain

import (
	"errors"
	"fmt"
	"log"
	"time"
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
	userNameInRoom, err := r.idToUserName(id)
	if err == nil {
		return fmt.Errorf("%w : '%v', '%v', '%v'", errID, id, r.Name, userNameInRoom)
	}

	if r.userNameInRoom(userName) {
		return fmt.Errorf("%w : '%v', '%v'", errUname, userName, r.Name)
	}

	r.UserList[userName] = id

	log.Printf("user '%s' connected to the room '%s'", userName, r.Name)

	return nil
}

// Leaver deletes user from the desired room.
func (r *Room) Leaver(userID string) (string, error) {
	userName, err := r.idToUserName(userID)
	if err != nil {
		return "", err
	}

	delete(r.UserList, userName)
	log.Printf("user '%s' disconnected from the room '%s'", userName, r.Name)

	return userName, nil
}

// Messenger gives server list of users in a room that have to receive given message.
func (r Room) Messenger(userID, roomName, text string) (Message, error) {
	m := Message{
		FromUserID: "",
		ToRoomName: roomName,
		ToID:       nil,
		Text:       text,
		Time:       time.Time{},
	}

	userName, err := r.idToUserName(userID)
	if err != nil {
		return m, err
	}

	m.FromUserID = userName

	for _, currentID := range r.UserList {
		if currentID == "SERVER" {
			continue
		}

		m.ToID = append(m.ToID, currentID)
	}

	return m, nil
}
