package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/holly-graham/schedulecli/cli"
	"github.com/holly-graham/schedulecli/db"
	"github.com/holly-graham/schedulecli/schedule"
)

const (
	addActivityCmd  = "Add Activity"
	viewScheduleCmd = "View Schedule"
	backCmd         = "Back"
	view            = "View Day"
	addAnother      = "Add to Same Day"
	mainMenu        = "Main Menu"
	weekOverview    = "Week Overview"
	dayOverview     = "Day Overview"
	differentDay    = "Add to a Different Day"
)

func main() {
	db, err := db.ConnectDatabase("activities_db.config")
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	scheduleService := schedule.NewService(db)

	cliMenu := cli.New(scheduleService)

	cliMenu.MainMenu()

}
