package main

import (
	"fmt"
	"lp3/internal/data"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// Displays all departments for uploading scans/reports
func (app *application) scanHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/scanCat.html")
	if err != nil {
		app.logger.Printf("Unable to get template File: scanCat.html: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}

// Diplays all scans/reports of a user for a given deparment
func (app *application) scanCatHandler(w http.ResponseWriter, r *http.Request) {
	cat := mux.Vars(r)["cat"]
	user, _ := app.getUserFromCookie(r)
	files, err := app.model.Fm.GetUserFilesByCategory(user.Id, cat, data.SCAN)
	if err != nil {
		app.logger.Printf("Unable to get files from database: %v\n", err)
		return
	}
	tmpl, err := template.ParseFiles("templates/scanView.html")
	if err != nil {
		app.logger.Printf("Unable to get template scanView.html: %v\n", err)
		return
	}
	tmpl.Execute(w, files)
}

// Gets all departments for uploading scans/reports
func (app *application) scanUploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/hello.html")
	fmt.Println(tmpl.ParseName)
	if err != nil {
		app.logger.Printf("Unable to get template File: hello.html: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}

// GET: Displays upload page
// POST: goes through all uploaded files and saves them to server as UUID.gz
func (app *application) scanCatUploadHandler(w http.ResponseWriter, r *http.Request) {
	// if err != nil {
	// 	http.Redirect(w, r, "", http.StatusFound)
	// }
	cat := mux.Vars(r)["cat"]
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/target.html")
		path := fmt.Sprintf("http://localhost:8000/scans/upload/%s", cat)
		fmt.Println(path)
		tmpl.Execute(w, path)
	} else {
		userId, _ := app.getUserFromCookie(r)
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			app.logger.Printf("Error while parsing multipart form: %v\n", err)
			return
		}

		files := r.MultipartForm.File["files[]"]
		for _, header := range files {
			file, err := header.Open()
			if err != nil {
				fmt.Printf("Unable to open file: %v\n", err)
				return
			}
			defer file.Close()

			f := data.NewFile(file, header, cat, userId.Id, data.SCAN)
			app.model.Fm.Insert(f)

			err = app.zip(f.Name, file)
			if err != nil {
				app.logger.Printf("Unable to zip file: %s, %v\n", err.Error(), err)
				return
			}
		}

		http.Redirect(w, r, "https://medilink.onrender.com/", http.StatusFound)
	}
}
