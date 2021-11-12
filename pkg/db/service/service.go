package dbservice

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	generaldb "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	rrep "github.com/nndergunov/RTGC-Project/pkg/db/repository/room"
	urep "github.com/nndergunov/RTGC-Project/pkg/db/repository/user"
)

type ServiceDB struct {
	db               *sql.DB
	roomsTable       *rrep.Rooms
	usersInRoomTable *urep.UsersInRoom
}

func (s *ServiceDB) Init(dbSource string) {
	database, err := generaldb.NewDB(dbSource)
	if err != nil {
		log.Fatalf("database open: %v", err)
	}

	s.db = database

	s.roomsTable = rrep.NewRoomsRepository(s.db)
	s.usersInRoomTable = urep.NewUsersInRoomRepository(s.db)
}

// ReadAllRooms gets all rooms from the database table.
func (s ServiceDB) ReadAllRooms() ([]*model.Rooms, error) {
	rooms, err := s.roomsTable.ListRooms(
		context.Background(),
		&generaldb.ListOptions{
			Sort: []generaldb.SortOrder{
				{
					Property: "id",
				},
			},
		},
		&generaldb.RoomsCriteria{
			ID:   nil,
			Name: nil,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("room list read: %w", err)
	}

	return rooms, nil
}

// ReadAllUsers gets all users from the database table.
func (s ServiceDB) ReadAllUsers() ([]*model.Usersinroom, error) {
	users, err := s.usersInRoomTable.ListUsersInRoom(
		context.Background(),
		&generaldb.ListOptions{
			Sort: []generaldb.SortOrder{
				{
					Property: "id",
				},
			},
		},
		&generaldb.UsersInRoomCriteria{
			ID:       nil,
			RoomID:   nil,
			UserID:   nil,
			Username: nil,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("users read list: %w", err)
	}

	return users, nil
}

// AddRoom creates new room instance in the database table.
func (s *ServiceDB) AddRoom(name string) (int, error) {
	id, err := s.roomsTable.CreateRooms(context.Background(), name)
	if err != nil {
		return 0, fmt.Errorf("rooms table write: %w", err)
	}

	return id, nil
}

// AddUser creates new user in the database table.
func (s *ServiceDB) AddUser(roomID int, userID, userName string) (int, error) {
	id, err := s.usersInRoomTable.CreateUsersInRoom(context.Background(), roomID, userID, userName)
	if err != nil {
		return 0, fmt.Errorf("users table write: %w", err)
	}

	return id, nil
}

// DelRoom deletes room from the database table.
func (s *ServiceDB) DelRoom(id int) error {
	err := s.roomsTable.DeleteRooms(context.Background(), id)
	if err != nil {
		return fmt.Errorf("room table delete: %w", err)
	}

	return nil
}

// DelUser removes user from the database table.
func (s *ServiceDB) DelUser(id int) error {
	err := s.usersInRoomTable.DeleteUsersInRoom(context.Background(), id)
	if err != nil {
		return fmt.Errorf("users table delete: %w", err)
	}

	return nil
}
