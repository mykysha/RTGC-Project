package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	dbservice "github.com/nndergunov/RTGC-Project/pkg/db/service"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
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

	db := dbservice.ServiceDB{}

	db.Init(dbSource)

	delUsers(db)
	delRooms(db)
}

func delRooms(db dbservice.ServiceDB) {
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

func delUsers(db dbservice.ServiceDB) {
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
