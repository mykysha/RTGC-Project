package app

import "log"

// Connecter adds the user to the desired room.
func Connecter(id, userName, roomName string) {
	if _, ok := RoomList[roomName]; !ok {
		NewRoom(userName, roomName)
	}

	r := RoomList[roomName]

	// TODO check if username is already taken.

	r.UserList[userName] = id

	log.Printf("\n"+"User %s connected to the room %s", userName, roomName)
}

func Leaver(userName, roomName string) {
	// TODO check if user is in the room

	r := RoomList[roomName]

	delete(r.UserList, userName)
}
