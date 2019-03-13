package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/holly-graham/scheduleapi/db"
	"github.com/holly-graham/scheduleapi/schedule"
	"github.com/holly-graham/scheduleapi/server"
)

const port = ":8000"

func main() {
	db, err := db.ConnectDatabase("activities_db.config")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	scheduleService := schedule.NewService(db)

	scheduleServer := server.NewServer(scheduleService)

	router := mux.NewRouter()
	router.HandleFunc("/day/{chosenDay}/activities", scheduleServer.ListActivitiesHandler).Methods("GET")
	router.HandleFunc("/day/{chosenDay}/activities", scheduleServer.AddActivityHandler).Methods("POST")

	http.Handle("/", router)

	fmt.Println("Waiting for requests on port:", port)
	http.ListenAndServe(port, nil)

}
