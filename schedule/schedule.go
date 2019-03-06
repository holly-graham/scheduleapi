package schedule

import "database/sql"

const (
	Monday    = "Monday"
	Tuesday   = "Tuesday"
	Wednesday = "Wednesday"
	Thursday  = "Thursday"
	Friday    = "Friday"
	Saturday  = "Saturday"
	Sunday    = "Sunday"
)

var (
	mondayActivities    []Activity
	tuesdayActivities   []Activity
	wednesdayActivities []Activity
	thursdayActivities  []Activity
	fridayActivities    []Activity
	saturdayActivities  []Activity
	sundayActivities    []Activity
)

type ScheduleService struct {
	db *sql.DB
}

type Activity struct {
	Time        string
	Description string
}

func NewService(db *sql.DB) *ScheduleService {
	return &ScheduleService{
		db: db,
	}
}

const (
	insertActivityQuery = "INSERT INTO activity (day, time, description) VALUES (?, ?, ?)"

	selectListQuery = "SELECT time, description FROM activity WHERE day = ?"
)

func (a Activity) String() string {
	return a.Time + ": " + a.Description
}

func (a *ScheduleService) AddActivity(activity Activity, day string) error {
	_, err := a.db.Exec(insertActivityQuery, day, activity.Time, activity.Description)

	if err != nil {
		return err
	}

	return nil
}

// func (a *ScheduleService) AddActivity(activity Activity, day string) {
// 	switch day {
// 	case Monday:
// 		mondayActivities = append(mondayActivities, activity)
// 	case Tuesday:
// 		tuesdayActivities = append(tuesdayActivities, activity)
// 	case Wednesday:
// 		wednesdayActivities = append(wednesdayActivities, activity)
// 	case Thursday:
// 		thursdayActivities = append(thursdayActivities, activity)
// 	case Friday:
// 		fridayActivities = append(fridayActivities, activity)
// 	case Saturday:
// 		saturdayActivities = append(saturdayActivities, activity)
// 	case Sunday:
// 		sundayActivities = append(sundayActivities, activity)
// 	}

// }

func (a *ScheduleService) ListActivities(day string) ([]Activity, error) {
	rows, err := a.db.Query(selectListQuery, day)
	if err != nil {
		return nil, err
	}

	var activities []Activity
	for rows.Next() {
		var activity Activity

		err := rows.Scan(
			&activity.Time,
			&activity.Description,
		)
		if err != nil {
			return nil, err
		}

		activities = append(activities, activity)
	}

	return activities, nil
}

// func (a *ScheduleService) ListActivities(day string) []Activity {
// 	switch day {
// 	case Monday:
// 		return mondayActivities
// 	case Tuesday:
// 		return tuesdayActivities
// 	case Wednesday:
// 		return wednesdayActivities
// 	case Thursday:
// 		return thursdayActivities
// 	case Friday:
// 		return fridayActivities
// 	case Saturday:
// 		return saturdayActivities
// 	case Sunday:
// 		return sundayActivities
// 	}

// 	return nil
// }
