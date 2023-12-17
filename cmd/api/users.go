package main

import (
	"database/sql"
	"errors"
	"fmt"
	"lp3/internal/data"
	"net/http"
	"text/template"
	"time"
)

func (app *application) registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	user := data.NewUser(username, password, email)
	if err := app.model.Um.Insert(user); err != nil {
		app.logger.Printf("Unable to insert user to DB: %v\n", err)
		return
	}

	http.Redirect(w, r, "/auth", http.StatusFound)
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := app.model.Um.GetUserByEmail(email)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			app.logger.Printf("Invalid Email Address: %v\n", err)
			http.Redirect(w, r, "../healthcheck", http.StatusFound)
		default:
			app.logger.Printf("Unable to retrieve user: %v\n", err)
			http.Redirect(w, r, "../healthcheck", http.StatusFound)
		}
	}
	if !user.CompareHashAndPassword(password) {
		app.logger.Println("Incorrect Password, Please try again")
		http.Redirect(w, r, "../healthcheck", http.StatusFound)
	}
	cookie := &http.Cookie{
		Name:  "user-cookie",
		Value: fmt.Sprintf("%d", user.Id),
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "../vitalsn", http.StatusFound)
}

// TODO: send email or sms to user's email or phone number, reminding them
func (app *application) reminderHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := app.getUserFromCookie(r)

	reminders, err := app.model.Rm.GetRemindersByUser(user.Id)
	if err != nil {
		app.logger.Printf("%v\n", err)
	}

	tmpl, err := template.ParseFiles("templates/reminders.html")
	if err != nil {
		app.logger.Printf("Unable to locate template file: %v\n", err)
		return
	}
	tmpl.Execute(w, reminders)

}

func (app *application) setReminderHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := app.getUserFromCookie(r)

	r.ParseForm()

	appointmentDate, _ := time.Parse("2006-01-02", r.FormValue("appointmentDate"))
	appointmentTime, _ := time.Parse("15:04", r.FormValue("appointmentTime"))
	hour, min, sec := appointmentTime.Clock()

	appointmentDateTime := appointmentDate.Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec))
	appointmentNote := r.FormValue("appointmentNote")
	email := r.FormValue("patientEmail")

	rem := data.NewReminder(appointmentDateTime, email, appointmentNote, user.Id)

	err := app.model.Rm.Insert(rem)
	if err != nil {
		app.logger.Printf("Unable to insert reminder: %v\n", err)
		return
	}

	http.Redirect(w, r, "../../", http.StatusFound)
}

func (app *application) vitalsNHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/vitals.html")
	if err != nil {
		app.logger.Printf("Unable to get template file: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}

func (app *application) vitalsYHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/vitalsx.html")
	if err != nil {
		app.logger.Printf("Unable to get template file: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}
