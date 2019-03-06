package cli

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/holly-graham/schedulecli/schedule"
	"github.com/manifoldco/promptui"
)

type CLI struct {
	scheduleService *schedule.ScheduleService
}

func New(scheduleService *schedule.ScheduleService) *CLI {
	return &CLI{
		scheduleService: scheduleService,
	}
}

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

func (c *CLI) MainMenu() {

	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				addActivityCmd,
				viewScheduleCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {

		case addActivityCmd:
			err := c.addActivityPrompt()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

		case viewScheduleCmd:
			err := c.viewSchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

}

func (c *CLI) addActivityPrompt() error {

	dayPrompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			schedule.Monday,
			schedule.Tuesday,
			schedule.Wednesday,
			schedule.Thursday,
			schedule.Friday,
			schedule.Saturday,
			schedule.Sunday,
			backCmd,
		},
	}

	_, chosenDay, err := dayPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch chosenDay {
	case backCmd:
		return nil
	default:
		err := c.addTimeDescription(chosenDay)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	}

	time.Sleep(500 * time.Millisecond)
	return nil
}

func (c *CLI) addTimeDescription(day string) error {
	for {
		timePrompt := promptui.Prompt{
			Label: "Time",
		}
		time, err := timePrompt.Run()
		if err != nil {
			return err
		}

		Description := promptui.Prompt{
			Label: "Description",
		}

		description, err := Description.Run()
		if err != nil {
			return err
		}

		Activities := schedule.Activity{
			Time:        time,
			Description: description,
		}

		c.scheduleService.AddActivity(Activities, day)

		fmt.Println("Added to", day, time, description)

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				addAnother,
				view,
				differentDay,
				mainMenu,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		switch result {

		case view:
			err := c.viewGivenDay(day)
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}

		// viewDay, err := c.scheduleService.ListActivities(day)
		// if err != nil {
		// 	fmt.Printf("Prompt failed %v\n", err)
		// 	return err
		// }

		// for _, day := range viewDay {
		// 	fmt.Println(day)

		// 	//fmt.Println("here", len(viewDay))

		// 	return nil
		// }

		case addAnother:
			continue

		case differentDay:
			c.addActivityPrompt()

		case mainMenu:
			return nil
		}

	}

	return nil
}

func (c *CLI) viewSchedule() error {

	for {
		fmt.Println()

		prompt := promptui.Select{
			Label: "Select One",
			Items: []string{
				weekOverview,
				dayOverview,
				backCmd,
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		switch result {

		case weekOverview:
			err := c.weekSchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}

		case dayOverview:
			err := c.daySchedule()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return err
			}
		case backCmd:
			return nil
		}

		time.Sleep(500 * time.Millisecond)

	}
}

func (c *CLI) weekSchedule() error {

	week := []string{
		schedule.Monday,
		schedule.Tuesday,
		schedule.Wednesday,
		schedule.Thursday,
		schedule.Friday,
		schedule.Saturday,
		schedule.Sunday,
	}

	for _, day := range week {

		fmt.Println(day)
		err := c.viewGivenDay(day)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
		time.Sleep(250 * time.Millisecond)

	}

	finalMenu := promptui.Select{
		Label: "Select One",
		Items: []string{
			weekOverview,
			dayOverview,
			backCmd,
			mainMenu,
		},
	}

	_, result, err := finalMenu.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {

	case weekOverview:
		err := c.weekSchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case dayOverview:
		err := c.daySchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}
	case backCmd:
		return nil

	case mainMenu:
		c.MainMenu()
	}
	return nil
}

func (c *CLI) daySchedule() error {

	dayPrompt := promptui.Select{
		Label: "Select Day",
		Items: []string{
			schedule.Monday,
			schedule.Tuesday,
			schedule.Wednesday,
			schedule.Thursday,
			schedule.Friday,
			schedule.Saturday,
			schedule.Sunday,
			backCmd,
			mainMenu,
		},
	}

	_, result, err := dayPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {

	//addTimeDescription that works for each day of the week

	case schedule.Monday:
		err := c.viewGivenDay(schedule.Monday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Tuesday:
		err := c.viewGivenDay(schedule.Tuesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Wednesday:
		err := c.viewGivenDay(schedule.Wednesday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Thursday:
		err := c.viewGivenDay(schedule.Thursday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Friday:
		err := c.viewGivenDay(schedule.Friday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Saturday:
		err := c.viewGivenDay(schedule.Saturday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case schedule.Sunday:
		err := c.viewGivenDay(schedule.Sunday)
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case backCmd:
		return nil

	case mainMenu:
		c.MainMenu()

	}

	return nil
}

func (c *CLI) viewGivenDay(day string) error {

	viewDay, err := c.scheduleService.ListActivities(day)
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	for _, day := range viewDay {
		fmt.Println(day)
	}

	return nil
}

func (c *CLI) menuAfterSchedule(day string) error {

	c.viewGivenDay(day)

	finalMenu := promptui.Select{
		Label: "Select One",
		Items: []string{
			weekOverview,
			dayOverview,
			backCmd,
		},
	}

	_, result, err := finalMenu.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	switch result {
	case weekOverview:
		err := c.weekSchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case dayOverview:
		err := c.daySchedule()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

	case backCmd:
		return nil
	}

	return nil
}
