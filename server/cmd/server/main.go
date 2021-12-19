package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nndergunov/RTGC-Project/server/api"
	"github.com/nndergunov/RTGC-Project/server/cmd/server/config"
	"github.com/nndergunov/RTGC-Project/server/pkg/app"
	"github.com/nndergunov/RTGC-Project/server/pkg/app/allrooms"
	dbservice "github.com/nndergunov/RTGC-Project/server/pkg/db/service"

	"github.com/joho/godotenv"
)

func main() {
	db := &dbservice.ServiceDB{}

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
