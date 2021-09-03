package app

import "log"

// Connects the user to the desired room.
func Connecter(userName, roomName string) error {
	if _, ok := RoomList[roomName]; !ok {
		NewRoom(userName, roomName)
	}

	r := RoomList[roomName]

	r.UserList = append(r.UserList, userName)

	log.Printf("\n"+"User %s connected to the room %s", userName, roomName)

	return nil
}
