package allroomsservice

import (
	"fmt"
	"log"

	roomservice "github.com/nndergunov/RTGC-Project/pkg/app/room"
	database "github.com/nndergunov/RTGC-Project/pkg/db/service"
	"github.com/nndergunov/RTGC-Project/pkg/domain"
)

// ServerUserName exists as a user in every room.
const ServerUserName = "SERVER"

type AllRooms struct {
	rooms        map[string]*roomservice.RoomService
	roomNameToID map[string]int
	db           *database.Database
}

func (a *AllRooms) Init() {
	a.rooms = make(map[string]*roomservice.RoomService)
	a.roomNameToID = make(map[string]int)

	a.db = &database.Database{}
	a.db.Init()

	err := a.fetchRooms()
	if err != nil {
		log.Fatalf("database initial rooms read: %v", err)
	}

	err = a.fetchUsers()
	if err != nil {
		log.Fatalf("database initial users read: %v", err)
	}
}

// fetchRooms reads all rooms stored in the database.
func (a *AllRooms) fetchRooms() error {
	rooms, err := a.db.ReadAllRooms()
	if err != nil {
		return fmt.Errorf("database (rooms table): %w", err)
	}

	for _, val := range rooms {
		if !a.roomExists(*val.Name) {
			nr := roomservice.RoomService{
				Room: &domain.Room{
					Name: *val.Name,
				},
				UserList:      make(map[string]string),
				UserIDToRowID: make(map[string]int),
			}
			a.rooms[*val.Name] = &nr
			nr.UserList[ServerUserName] = ServerUserName

			a.roomNameToID[*val.Name] = int(val.ID)

			log.Printf("room %s created (from db)", *val.Name)
		}
	}

	return nil
}

// fetchUsers reads all users stored in the database.
func (a *AllRooms) fetchUsers() error {
	users, err := a.db.ReadAllUsers()
	if err != nil {
		return fmt.Errorf("database (users table): %w", err)
	}

	for _, val := range users {
		roomName, f := a.findRoomByID(int(*val.RoomID))
		if !f {
			log.Printf("room with key %d not found", val.ID)

			continue
		}

		newRoom := a.rooms[roomName]

		err := newRoom.Connecter(*val.Userid, *val.Username)
		if err != nil {
			return fmt.Errorf("joinHandler: %w", err)
		}

		log.Printf("user %s connected to the room %s (from db)", *val.Username, roomName)
	}

	return nil
}

// newRoom creates new room.
func (a *AllRooms) newRoom(userName, roomName string) error {
	nr := roomservice.RoomService{
		Room: &domain.Room{
			Name: roomName,
		},
		UserList:      make(map[string]string),
		UserIDToRowID: make(map[string]int),
	}
	a.rooms[roomName] = &nr
	nr.UserList[ServerUserName] = ServerUserName

	id, err := a.db.AddRoom(roomName)
	if err != nil {
		return fmt.Errorf("database: %w", err)
	}

	a.roomNameToID[roomName] = id

	log.Printf("user '%s' created new room '%s'", userName, roomName)

	return nil
}

func (a AllRooms) roomExists(roomName string) bool {
	if _, ok := a.rooms[roomName]; ok {
		return true
	}

	return false
}

func (a AllRooms) findRoomByID(id int) (string, bool) {
	for key, val := range a.roomNameToID {
		if val == id {
			return key, true
		}
	}

	return "", false
}

// Join connects user to the desired room.
func (a *AllRooms) Join(userID, userName, roomName string) error {
	if userName == "SERVER" || userName == "ADMIN" {
		return fmt.Errorf("%w: '%s'", errUnsupportedUsername, userName)
	}

	if !a.roomExists(roomName) {
		err := a.newRoom(userName, roomName)
		if err != nil {
			return fmt.Errorf("join: %w", err)
		}
	}

	newRoom := a.rooms[roomName]

	err := newRoom.Connecter(userID, userName)
	if err != nil {
		return fmt.Errorf("joinHandler: %w", err)
	}

	id, err := a.db.AddUser(a.roomNameToID[roomName], userID, userName)
	if err != nil {
		return fmt.Errorf("database add: %w", err)
	}

	newRoom.UserIDToRowID[userID] = id

	return nil
}

// Leave disconnects user from desired room.
func (a *AllRooms) Leave(userID, roomName, text string) (string, error) {
	if text != "-" {
		log.Printf("'%s' reason to leave from '%s': '%s'", userID, roomName, text)
	}

	if !a.roomExists(roomName) {
		return "", fmt.Errorf("%w: '%s'", errNoRoom, roomName)
	}

	newRoom := a.rooms[roomName]

	uName, err := newRoom.Leaver(userID)
	if err != nil {
		return "", fmt.Errorf("leaveHandler: %w", err)
	}

	err = a.db.DelUser(newRoom.UserIDToRowID[userID])
	if err != nil {
		err = fmt.Errorf("database: %w", err)
	}

	return uName, err
}

// Send delivers message to the desired room.
func (a AllRooms) Send(id, roomName, text string) (*domain.Message, error) {
	if !a.roomExists(roomName) {
		return nil, fmt.Errorf("%w: '%s'", errNoRoom, roomName)
	}

	newRoom := a.rooms[roomName]

	m, err := newRoom.Messenger(id, roomName, text)
	if err != nil {
		err = fmt.Errorf("send: %w", err)

		return nil, err
	}

	return m, nil
}
