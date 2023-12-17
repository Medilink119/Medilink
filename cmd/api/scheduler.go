package main

import (
	"fmt"
	"time"
)

func (app *application) runScheduler() {
	app.scheduler.Every(1).Hour().Do(func() {
		app.logger.Println("Running Scheduler....")
		app.checkAppoitnments()
	})
	app.scheduler.StartBlocking()
}

func (app *application) checkAppoitnments() {
	reminders, err := app.model.Rm.GetAllUnsentReminders()
	if err != nil {
		app.logger.Printf("Unable to get reminders: %v\n", err)
		return
	}
	for _, reminder := range reminders {
		if time.Until(app.ConvertToIST(reminder.Appointment)) < time.Hour {
			fmt.Println(reminder.Email)
			err := app.mailer.Send(reminder.Email)
			if err != nil {
				app.logger.Printf("Unable to send mail: %v\n", err)
				continue
			}
			app.model.Rm.UpdateReminderSent(reminder.Id)
		}
	}
}
