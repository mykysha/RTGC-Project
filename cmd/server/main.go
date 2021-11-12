package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nndergunov/RTGC-Project/api"
	"github.com/nndergunov/RTGC-Project/cmd/server/config"
	"github.com/nndergunov/RTGC-Project/pkg/app"
	"github.com/nndergunov/RTGC-Project/pkg/app/allrooms"
	dbservice "github.com/nndergunov/RTGC-Project/pkg/db/service"
)

func main() {
	db := &dbservice.ServiceDB{}

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

	db.Init(dbSource)

	roomList := &allrooms.AllRooms{}
	roomList.Init(db)

	r := &app.Router{}

	r.Init(roomList)

	mux := config.MainServer()
	logger := log.New(os.Stdout, "server ", log.LstdFlags)

	a := &api.API{}
	a.Init(mux, logger, r)

	log.Fatal(http.ListenAndServe(":8080", a))
}
