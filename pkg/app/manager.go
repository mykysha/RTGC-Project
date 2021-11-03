package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	jet "github.com/nndergunov/RTGC-Project/pkg/db"
	"github.com/nndergunov/RTGC-Project/pkg/db/repository"
	dom "github.com/nndergunov/RTGC-Project/pkg/domain"
)

const SERVER = "SERVER"

type Router struct {
	roomList         map[string]*dom.Room
	roomNameToID     map[string]int
	myDB             *sql.DB
	roomsTable       *jet.RoomsRepository
	usersInRoomTable *jet.UsersInRoomRepository
}

func (r *Router) Init() {
	r.roomList = make(map[string]*dom.Room)
	r.roomNameToID = make(map[string]int)
	dbSource := fmt.Sprintf(
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + pass +
			" dbname=" + dbname +
			" sslmode=" + ssl,
	)

	database, err := jet.NewDB(dbSource)
	if err != nil {
		log.Fatalf("database open: %v", err)
	}

	r.myDB = database

	r.roomsTable = jet.NewRoomsRepository(r.myDB)
	r.usersInRoomTable = jet.NewUsersInRoomRepository(r.myDB)

	err = r.fetchRooms()
	if err != nil {
		log.Fatalf("database initial rooms read: %v", err)
	}

	err = r.fetchUsers()
	if err != nil {
		log.Fatalf("database initial users read: %v", err)
	}
}

// fetchRooms reads all rooms stored in the database.
func (r *Router) fetchRooms() error {
	roomListDB, err := r.roomsTable.ListRooms(
		context.Background(),
		&repository.ListOptions{
			Sort: []repository.SortOrder{
				{
					Property: "id",
				},
			},
		},
		&repository.RoomsCriteria{
			ID:   nil,
			Name: nil,
		},
	)
	if err != nil {
		return fmt.Errorf("rooms list: %w", err)
	}

	for _, val := range roomListDB {
		if !r.roomExists(*val.Name) {
			nr := dom.Room{
				Name:          *val.Name,
				UserList:      make(map[string]string),
				UserIDToRowID: make(map[string]int),
			}
			r.roomList[*val.Name] = &nr
			nr.UserList[SERVER] = SERVER

			r.roomNameToID[*val.Name] = int(val.ID)

			log.Printf("room %s created (from db)", *val.Name)
		}
	}

	return nil
}

// fetchUsers reads all users stored in the database.
func (r *Router) fetchUsers() error {
	userListDB, err := r.usersInRoomTable.ListUsersInRoom(
		context.Background(),
		&repository.ListOptions{
			Sort: []repository.SortOrder{
				{
					Property: "id",
				},
			},
		},
		&repository.UsersInRoomCriteria{
			ID:       nil,
			RoomID:   nil,
			UserID:   nil,
			Username: nil,
		},
	)
	if err != nil {
		return fmt.Errorf("user list: %w", err)
	}

	for _, val := range userListDB {
		roomName, f := r.findByID(int(*val.RoomID))
		if !f {
			log.Printf("room with key %d not found", val.ID)

			continue
		}

		room := r.roomList[roomName]

		err := room.Connecter(*val.Userid, *val.Username)
		if err != nil {
			return fmt.Errorf("joinHandler: %w", err)
		}

		log.Printf("user %s connected to the room %s (from db)", *val.Username, roomName)
	}

	return nil
}

// newRoom creates new room.
func (r *Router) newRoom(userName, roomName string) error {
	nr := dom.Room{
		Name:          roomName,
		UserList:      make(map[string]string),
		UserIDToRowID: make(map[string]int),
	}
	r.roomList[roomName] = &nr
	nr.UserList[SERVER] = SERVER

	id, err := r.roomsTable.CreateRooms(context.Background(), roomName)
	if err != nil {
		return fmt.Errorf("database write: %w", err)
	}

	r.roomNameToID[roomName] = id

	log.Printf("user '%s' created new room '%s'", userName, roomName)

	return nil
}

func (r Router) roomExists(roomName string) bool {
	if _, ok := r.roomList[roomName]; ok {
		return true
	}

	return false
}

func (r Router) findByID(id int) (string, bool) {
	for key, val := range r.roomNameToID {
		if val == id {
			return key, true
		}
	}

	return "", false
}
