package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) route() *mux.Router {
	router := mux.NewRouter()

	// Landing Page Handler
	router.HandleFunc("/", app.indexPageHandler)

	// User Authentication Pages
	router.HandleFunc("/auth", app.authHandler)
	router.HandleFunc("/login", app.loginHandler).Methods("POST")
	router.HandleFunc("/register", app.registerHandler).Methods("POST")

	// Displays File given it's ID
	router.HandleFunc("/file/{id}", app.fileViewer)

	// Prescrition Pages
	router.HandleFunc("/prescriptions", app.checkUser(app.prescriptionHandler))
	router.HandleFunc("/prescriptions/upload", app.checkUser(app.prescriptionUploadHandler))
	router.HandleFunc("/prescriptions/file/{id}", app.fileViewer)

	// Scan Pages
	// Displays Scan Deparments
	router.HandleFunc("/scans", app.checkUser(app.scanHandler))
	// Displays Files for a given deparment
	router.HandleFunc("/scans/view/{cat}", app.checkUser(app.scanCatHandler))
	// Displays Scan Departments for uploading
	router.HandleFunc("/scans/upload", app.checkUser(app.scanUploadHandler))
	// Upload page for scans/reports
	router.HandleFunc("/scans/upload/{cat}", app.checkUser(app.scanCatUploadHandler))

	// reminder pages
	router.HandleFunc("/reminders", app.reminderHandler)
	router.HandleFunc("/reminders/set", app.setReminderHandler)

	// Vitals page
	router.HandleFunc("/vitalsn", app.vitalsNHandler)
	router.HandleFunc("/vitalsy", app.vitalsYHandler)

	// Health Check
	router.HandleFunc("/healthcheck", app.healthcheckHandler)

	router.HandleFunc("/tryupload", app.tryUploadHandler)
	router.HandleFunc("/try", app.tryHandler)

	// Holds Static Files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	return router
}
