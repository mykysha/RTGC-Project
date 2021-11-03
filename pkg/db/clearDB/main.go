package main

import (
	"context"
	"fmt"
	"log"

	jet "github.com/nndergunov/RTGC-Project/pkg/db"
	"github.com/nndergunov/RTGC-Project/pkg/db/repository"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "rtgc"
	pass := "rtgcpass"
	dbname := "rtgc"
	ssl := "disable"
	dataSourceName := fmt.Sprintf(
		"host=" + host +
			" port=" + port +
			" user=" + user +
			" password=" + pass +
			" dbname=" + dbname +
			" sslmode=" + ssl,
	)

	myDB, err := jet.NewDB(dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	roomsTable := jet.NewRoomsRepository(myDB)
	usersInRoomTable := jet.NewUsersInRoomRepository(myDB)

	delUsers(usersInRoomTable)
	delRooms(roomsTable)
}

func delRooms(roomsTable *jet.RoomsRepository) {
	roomListDB, err := roomsTable.ListRooms(
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
		log.Fatal(err)
	}

	for _, val := range roomListDB {
		err = roomsTable.DeleteRooms(context.Background(), int(val.ID))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func delUsers(usersInRoomTable *jet.UsersInRoomRepository) {
	userListDB, err := usersInRoomTable.ListUsersInRoom(
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
		log.Fatal(err)
	}

	for _, val := range userListDB {
		err = usersInRoomTable.DeleteUsersInRoom(context.Background(), int(val.ID))
		if err != nil {
			log.Fatal(err)
		}
	}
}
