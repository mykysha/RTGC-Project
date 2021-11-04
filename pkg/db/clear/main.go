package main

import (
	"log"

	database "github.com/nndergunov/RTGC-Project/pkg/db/service"
)

func main() {
	db := database.Database{}
	db.Init()

	delUsers(db)
	delRooms(db)
}

func delRooms(db database.Database) {
	rooms, err := db.ReadAllRooms()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range rooms {
		err = db.DelRoom(int(val.ID))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func delUsers(db database.Database) {
	users, err := db.ReadAllUsers()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range users {
		err = db.DelUser(int(val.ID))
		if err != nil {
			log.Fatal(err)
		}
	}
}
