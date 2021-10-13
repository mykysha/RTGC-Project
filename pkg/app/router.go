package app

import (
	"errors"
	"fmt"
	"log"
)

// static errors.
var (
	errUnknownAction       = errors.New("action not supported")
	errUnsupportedUsername = errors.New("username is not supported")
	errNoRoom              = errors.New("no room with such name exists")
)

// ActionHandler sends request to the correct handler.
func (r Router) ActionHandler(id, action, roomName, userName, text string) (string, string, string, []string, error) {
	switch action {
	case "join":
		err := r.joinHandler(id, userName, roomName)
		if err != nil {
			return "", "", "", nil, err
		}

		joinMessage := fmt.Sprintf("user '%s' joined the room '%s'", userName, roomName)

		return r.ActionHandler("SERVER", "send", roomName, "SERVER", joinMessage)

	case "leave":
		userName, err := r.leaveHandler(id, roomName, text)
		if err != nil {
			return "", "", "", nil, err
		}

		leaveMessage := fmt.Sprintf("user '%s' left the room '%s'", userName, roomName)

		return r.ActionHandler("SERVER", "send", roomName, "SERVER", leaveMessage)

	case "send":
		return r.sendHandler(id, roomName, text)

	default:
		return "", "", "", nil, fmt.Errorf("%w : '%s'", errUnknownAction, action)
	}
}

// joinHandler routes join request to the desired room.
func (r Router) joinHandler(id, userName, roomName string) error {
	if userName == "SERVER" || userName == "ADMIN" {
		return fmt.Errorf("%w : '%s'", errUnsupportedUsername, userName)
	}

	if !r.roomExists(roomName) {
		r.newRoom(userName, roomName)
	}

	room := r.roomList[roomName]

	err := room.Connecter(id, userName)
	if err != nil {
		return fmt.Errorf("joinHandler: %w", err)
	}

	return nil
}

// leaveHandler routes leave request to the desired room.
func (r Router) leaveHandler(id, roomName, text string) (string, error) {
	if text != "-" {
		log.Printf("'%s' reason to leave from '%s': '%s'", id, roomName, text)
	}

	if !r.roomExists(roomName) {
		return "", fmt.Errorf("%w : '%s'", errNoRoom, roomName)
	}

	room := r.roomList[roomName]

	uName, err := room.Leaver(id)
	if err != nil {
		err = fmt.Errorf("leaveHandler: %w", err)
	}

	return uName, err
}

// sendHandler routes send request to the desired room.
func (r Router) sendHandler(id, roomName, text string) (string, string, string, []string, error) {
	if !r.roomExists(roomName) {
		return "", "", "", nil, fmt.Errorf("%w : '%s'", errNoRoom, roomName)
	}

	room := r.roomList[roomName]

	m, err := room.Messenger(id, roomName, text)
	if err != nil {
		err = fmt.Errorf("send: %w", err)

		return "", "", "", nil, err
	}

	return m.FromUserID, m.ToRoomName, m.Text, m.ToID, nil
}
