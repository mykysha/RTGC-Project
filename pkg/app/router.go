package app

import (
	"fmt"
	"log"
)

// ActionHandler sends request to the correct handler.
func ActionHandler(id, action, roomName, userName, text string) (string, string, string, []string, error) {
	switch action {
	case "join":
		joinErr := joinHandler(id, userName, roomName)

		if joinErr != nil {
			return "", "", "", nil, joinErr
		}

		joinMessage := fmt.Sprintf("user '%s' joined the room '%s'", userName, roomName)

		return ActionHandler("SERVER", "send", roomName, "SERVER", joinMessage)

	case "leave":
		userName, leaveErr := leaveHandler(id, roomName, text)

		if leaveErr != nil {
			return "", "", "", nil, leaveErr
		}

		leaveMessage := fmt.Sprintf("user '%s' left the room '%s'", userName, roomName)

		return ActionHandler("SERVER", "send", roomName, "SERVER", leaveMessage)

	case "send":
		return sendHandler(id, roomName, text)

	default:
		unknownAction := fmt.Errorf("action '%s' not supported", action)

		return "", "", "", nil, unknownAction
	}
}

// joinHandler routes join request to the desired room.
func joinHandler(id, userName, roomName string) error {
	if userName == "SERVER" || userName == "ADMIN" {
		unsupportedNameError := fmt.Errorf("username '%s' is not supported", userName)

		return unsupportedNameError
	}

	if _, ok := roomList[roomName]; !ok {
		NewRoom(userName, roomName)
	}

	room := roomList[roomName]

	conErr := room.Connecter(id, userName)

	return conErr
}

// leaveHandler routes leave request to the desired room.
func leaveHandler(id, roomName, text string) (string, error) {
	log.Printf("\n"+"ID: '%s', Action: 'leave', RoomName: '%s'", id, roomName)

	if text != "-" {
		log.Printf("'%s' reason to leave: '%s'", id, text)
	}

	if !RoomExists(roomName) {
		errNoRoom := fmt.Errorf("found no room named '%s'", roomName)

		return "", errNoRoom
	}

	room := roomList[roomName]

	return room.Leaver(id)
}

// sendHandler routes send request to the desired room.
func sendHandler(id, roomName, text string) (string, string, string, []string, error) {
	if !RoomExists(roomName) {
		errNoRoom := fmt.Errorf("found no room named '%s'", roomName)

		return "", "", "", nil, errNoRoom
	}

	room := roomList[roomName]

	return room.Messenger(id, roomName, text)
}
