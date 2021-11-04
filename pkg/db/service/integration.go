package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	generaldb "github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public"
	"github.com/nndergunov/RTGC-Project/pkg/db/internal/rtgc/public/model"
	rrep "github.com/nndergunov/RTGC-Project/pkg/db/roomrepository"
	urep "github.com/nndergunov/RTGC-Project/pkg/db/usersrepository"
)

type Database struct {
	db               *sql.DB
	roomsTable       *rrep.RoomsRepository
	usersInRoomTable *urep.UsersInRoomRepository
}

func (d *Database) Init() {
	if err := godotenv.Load("db.env"); err != nil {
		log.Printf("env file read: %v", err)
	}

	dbSource := fmt.Sprintf(
		"host=" + os.Getenv("HOST") +
			" port=" + os.Getenv("PORT") +
			" user=" + os.Getenv("USER") +
			" password=" + os.Getenv("PASS") +
			" dbname=" + os.Getenv("NAME") +
			" sslmode=" + os.Getenv("SSL"),
	)

	database, err := generaldb.NewDB(dbSource)
	if err != nil {
		log.Fatalf("database open: %v", err)
	}

	d.db = database

	d.roomsTable = rrep.NewRoomsRepository(d.db)
	d.usersInRoomTable = urep.NewUsersInRoomRepository(d.db)
}

// ReadAllRooms gets all rooms from the database table.
func (d Database) ReadAllRooms() ([]*model.Rooms, error) {
	rooms, err := d.roomsTable.ListRooms(
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
func (d Database) ReadAllUsers() ([]*model.Usersinroom, error) {
	users, err := d.usersInRoomTable.ListUsersInRoom(
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
func (d *Database) AddRoom(name string) (int, error) {
	id, err := d.roomsTable.CreateRooms(context.Background(), name)
	if err != nil {
		return 0, fmt.Errorf("rooms table write: %w", err)
	}

	return id, nil
}

// AddUser creates new user in the database table.
func (d *Database) AddUser(roomID int, userID, userName string) (int, error) {
	id, err := d.usersInRoomTable.CreateUsersInRoom(context.Background(), roomID, userID, userName)
	if err != nil {
		return 0, fmt.Errorf("users table write: %w", err)
	}

	return id, nil
}

// DelRoom deletes room from the database table.
func (d *Database) DelRoom(id int) error {
	err := d.roomsTable.DeleteRooms(context.Background(), id)
	if err != nil {
		return fmt.Errorf("room table delete: %w", err)
	}

	return nil
}

// DelUser removes user from the database table.
func (d *Database) DelUser(id int) error {
	err := d.usersInRoomTable.DeleteUsersInRoom(context.Background(), id)
	if err != nil {
		return fmt.Errorf("users table delete: %w", err)
	}

	return nil
}
