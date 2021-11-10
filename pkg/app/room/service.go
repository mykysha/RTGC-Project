package room

import (
	"fmt"
	"log"
	"time"

	"github.com/nndergunov/RTGC-Project/pkg/domain"
)

// ServerUserName exists as a user in every room.
const ServerUserName = "SERVER"

type Room struct {
	Room          *domain.Room
	UserList      map[string]string // username - userid
	UserIDToRowID map[string]int    // userid - database row id
}

// Connecter adds the user to the desired room.
func (r *Room) Connecter(id, userName string) error {
	userNameInRoom, err := r.idToUserName(id)
	if err == nil {
		return fmt.Errorf("%w : '%v', '%v', '%v'", errID, id, r.Room.Name, userNameInRoom)
	}

	if r.userNameInRoom(userName) {
		return fmt.Errorf("%w : '%v', '%v'", errUname, userName, r.Room.Name)
	}

	r.UserList[userName] = id

	log.Printf("user '%s' connected to the room '%s' (id '%s')", userName, r.Room.Name, id)

	return nil
}

// Leaver deletes user from the desired room.
func (r *Room) Leaver(userID string) (string, error) {
	userName, err := r.idToUserName(userID)
	if err != nil {
		return "", err
	}

	delete(r.UserList, userName)
	log.Printf("user '%s' disconnected from the room '%s'", userName, r.Room.Name)

	return userName, nil
}

// Messenger gives server list of users in a room that have to receive given message.
func (r Room) Messenger(userID, roomName, text string) (*domain.Message, error) {
	m := domain.Message{
		FromUserName: "",
		ToRoomName:   roomName,
		ToID:         nil,
		Text:         text,
		Time:         time.Time{},
	}

	userName, err := r.idToUserName(userID)
	if err != nil {
		return nil, err
	}

	m.FromUserName = userName

	for _, currentID := range r.UserList {
		if currentID == ServerUserName {
			continue
		}

		m.ToID = append(m.ToID, currentID)
	}

	return &m, nil
}

func (r Room) idToUserName(userID string) (string, error) {
	var userName string

	for currentName, currentID := range r.UserList {
		if currentID == userID {
			userName = currentName

			return userName, nil
		}
	}

	return "", fmt.Errorf("%w : '%s', '%s'", errNoUser, userID, r.Room.Name)
}

func (r Room) userNameInRoom(userName string) bool {
	if _, ok := r.UserList[userName]; ok {
		return true
	}

	return false
}
