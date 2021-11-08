package db_test

import (
	"testing"

	database "github.com/nndergunov/RTGC-Project/pkg/db/service"
)

// Warning! Use with an empty database!

type testInfo struct {
	userTableID int
	roomTableID int
	userID      string
	userName    string
	roomName    string
	db          *database.Database
	test        *testing.T
}

func TestDatabase(t *testing.T) {
	t.Parallel()

	test := testInfo{
		userTableID: 0,
		roomTableID: 0,
		userID:      "testID1",
		userName:    "testName1",
		roomName:    "testRoom1",
		db:          nil,
		test:        t,
	}

	test.db = &database.Database{}
	test.db.Init()

	test.addNewRoom()
	test.addExistentRoom()
	test.addNewUser()
	test.addExistentUser()
	test.addUserToNonExistentRoom()
	test.readUsers()
	test.readRooms()
	test.delRoomWithUser()
	test.delUser()
	test.delRoom()
}

func (info *testInfo) addNewRoom() {
	roomTableID, err := info.db.AddRoom(info.roomName)
	if err != nil {
		info.test.Fatalf("Creating room: expected no err, got %v", err)
	}

	info.roomTableID = roomTableID
}

func (info *testInfo) addExistentRoom() {
	_, err := info.db.AddRoom(info.roomName)
	if err == nil {
		info.test.Fatalf("Creating existent room: expected err, got none")
	}
}

func (info *testInfo) addNewUser() {
	userTableID, err := info.db.AddUser(info.roomTableID, info.userID, info.userName)
	if err != nil {
		info.test.Fatalf("Creating user in existent room: expected no err, got %v", err)
	}

	info.userTableID = userTableID
}

func (info *testInfo) addExistentUser() {
	_, err := info.db.AddUser(info.roomTableID, info.userID, info.userName)
	if err == nil {
		info.test.Fatalf("Creating existent user: expected err, got none")
	}
}

func (info *testInfo) addUserToNonExistentRoom() {
	_, err := info.db.AddUser(info.roomTableID+1, info.userID+"2", info.userName+"2")
	if err == nil {
		info.test.Fatalf("Creating user in non-existent room: expected err, got none")
	}
}

func (info *testInfo) readUsers() {
	users, err := info.db.ReadAllUsers()
	if err != nil {
		info.test.Fatalf("Reading users: expected no err, got %v", err)
	}

	if len(users) > 1 {
		info.test.Fatalf("Reading users: expected 1 user, got %d", len(users))
	}

	for _, val := range users {
		if int(val.ID) != info.userTableID {
			info.test.Fatalf("Reading user table ID: expected %d, got %d", info.userTableID, val.ID)
		}

		if *val.Username != info.userName {
			info.test.Fatalf("Reading user name: expected %s, got %s", info.userName, *val.Username)
		}

		if *val.Userid != info.userID {
			info.test.Fatalf("Reading user ID: expected %s, got %s", info.userID, *val.Userid)
		}

		if int(*val.RoomID) != info.roomTableID {
			info.test.Fatalf("Reading room table ID from user: expected %d, got %d", info.roomTableID, *val.RoomID)
		}
	}
}

func (info *testInfo) readRooms() {
	rooms, err := info.db.ReadAllRooms()
	if err != nil {
		info.test.Fatalf("Reading rooms: expected no err, got %v", err)
	}

	if len(rooms) > 1 {
		info.test.Fatalf("Reading rooms: expected 1 room, got %d", len(rooms))
	}

	for _, val := range rooms {
		if int(val.ID) != info.roomTableID {
			info.test.Fatalf("Reading room table ID: expected %d, got %d", info.roomTableID, val.ID)
		}

		if *val.Name != info.roomName {
			info.test.Fatalf("Reading room name: expected %s, got %s", info.roomName, *val.Name)
		}
	}
}

func (info *testInfo) delRoomWithUser() {
	err := info.db.DelRoom(info.roomTableID)
	if err == nil {
		info.test.Fatal("Deleting room with user: expected err, got none")
	}
}

func (info *testInfo) delUser() {
	err := info.db.DelUser(info.userTableID)
	if err != nil {
		info.test.Fatalf("Deleting user: expected no err, got %v", err)
	}
}

func (info *testInfo) delRoom() {
	err := info.db.DelRoom(info.roomTableID)
	if err != nil {
		info.test.Fatalf("Deleting room: expected no err, got %v", err)
	}
}
