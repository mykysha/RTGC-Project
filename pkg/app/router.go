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
func ActionHandler(id, action, roomName, userName, text string) (string, string, string, []string, error) {
	switch action {
	case "join":
		err := joinHandler(id, userName, roomName)
		if err != nil {
			return "", "", "", nil, err
		}

		joinMessage := fmt.Sprintf("user '%s' joined the room '%s'", userName, roomName)

		return ActionHandler("SERVER", "send", roomName, "SERVER", joinMessage)

	case "leave":
		userName, err := leaveHandler(id, roomName, text)
		if err != nil {
			return "", "", "", nil, err
		}

		leaveMessage := fmt.Sprintf("user '%s' left the room '%s'", userName, roomName)

		return ActionHandler("SERVER", "send", roomName, "SERVER", leaveMessage)

	case "send":
		return sendHandler(id, roomName, text)

	default:
		return "", "", "", nil, fmt.Errorf("%w : '%s'", errUnknownAction, action)
	}
}

// joinHandler routes join request to the desired room.
func joinHandler(id, userName, roomName string) error {
	if userName == "SERVER" || userName == "ADMIN" {
		return fmt.Errorf("%w : '%s'", errUnsupportedUsername, userName)
	}

	if _, ok := roomList[roomName]; !ok {
		NewRoom(userName, roomName)
	}

	room := roomList[roomName]

	err := room.Connecter(id, userName)
	if err != nil {
		return fmt.Errorf("joinHandler: %w", err)
	}

	return nil
}

// leaveHandler routes leave request to the desired room.
func leaveHandler(id, roomName, text string) (string, error) {
	log.Printf("\n"+"ID: '%s', Action: 'leave', RoomName: '%s'", id, roomName)

	if text != "-" {
		log.Printf("'%s' reason to leave: '%s'", id, text)
	}

	if !RoomExists(roomName) {
		return "", fmt.Errorf("%w : '%s'", errNoRoom, roomName)
	}

	room := roomList[roomName]

	uName, err := room.Leaver(id)
	if err != nil {
		err = fmt.Errorf("leaveHandler: %w", err)
	}

	return uName, err
}

// sendHandler routes send request to the desired room.
func sendHandler(id, roomName, text string) (string, string, string, []string, error) {
	if !RoomExists(roomName) {
		return "", "", "", nil, fmt.Errorf("%w : '%s'", errNoRoom, roomName)
	}

	room := roomList[roomName]

	fromUser, fromRoom, message, toID, err := room.Messenger(id, roomName, text)
	if err != nil {
		err = fmt.Errorf("send: %w", err)

		return "", "", "", nil, err
	}

	return fromUser, fromRoom, message, toID, nil
}
