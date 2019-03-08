package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/holly-graham/scheduleapi/schedule"
)

type ScheduleServer struct {
	scheduleService *schedule.ScheduleService
}

func NewServer(scheduleService *schedule.ScheduleService) *ScheduleServer {
	return &ScheduleServer{
		scheduleService: scheduleService,
	}
}

func (s *ScheduleServer) ListActivitiesHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chosenDay := vars["chosenDay"]

	activities, err := s.scheduleService.ListActivities(chosenDay)
	if err != nil {
		fmt.Println("Error listing activities:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	activitiesJSON, err := json.Marshal(activities)
	if err != nil {
		fmt.Println("Error marshaling activities:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	rw.Write(activitiesJSON)
}

type CreateActivityRequest struct {
	Time        string
	Description int
}

func (s *ScheduleServer) AddActivityHandler(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chosenDay := vars["chosenDay"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body:", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	newScheduleActivity := schedule.Activity{
		Time:        newActivity.time,
		Description: newActivity.description,
	}

	s.scheduleService.AddActivity(newScheduleActivity, chosenDay)

	var newActivity CreateActivityRequest
	err = json.Unmarshal(requestBody, &newActivity)
	if err != nil {
		fmt.Println("Error unmarshaling new activity details:", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.scheduleService.AddActivity(newScheduleActivity, chosenDay)
	if err != nil {
		fmt.Println("Error creating activity:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
