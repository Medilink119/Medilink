package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := app.getUserFromCookie(r)
	if err != nil {
		app.logger.Printf("Unable to retrieve user from cookie: %v\n", err)
		return
	}
	healthData := &map[string]string{
		"Status":      "Available",
		"Environment": app.cfg.env,
	}
	if err := app.writeJSON(w, healthData); err != nil {
		app.logger.Printf("Could not convert health data to JSON: %v\n", err)
		return
	}
}

func (app *application) indexPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/buildd.html")
	if err != nil {
		app.logger.Printf("Unable to locate template file: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}

// Login/signup handler -> rename it to something better
func (app *application) authHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/login.html")
	if err != nil {
		app.logger.Printf("Unable to locate template file: %v\n", err)
		return
	}
	tmpl.Execute(w, nil)
}

// Middleware function, checks if the user id is present in the cookie
func (app *application) checkUser(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := app.getUserFromCookie(r)
		if err != nil {
			http.Redirect(w, r, "http://localhost:8000/auth", http.StatusFound)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Read file from disk and write to browser
func (app *application) fileViewer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fileId, _ := strconv.Atoi(id)

	file, err := app.model.Fm.GetFileById(fileId)
	if err != nil {
		app.logger.Printf("Unable to get file by id: %v\n", err)
		return
	}

	bytes, err := app.unzip(file.Name)
	if err != nil {
		app.logger.Printf("Unable to unzip file: %v\n", err)
		return
	}

	contentType := http.DetectContentType(bytes)

	w.Header().Set("Content-Type", contentType)
	w.Write(bytes)
}

// Create a zip file in uploads directory as UUID.gz
func (app *application) zip(fileName string, formFile multipart.File) error {
	dst, err := os.Create(fmt.Sprintf("./uploads/%s.gz", fileName))
	if err != nil {
		return errors.New("unable to create destination file")
	}
	defer dst.Close()

	fileBytes, err := io.ReadAll(formFile)
	if err != nil {
		return errors.New("unable to read form file")
	}

	w := gzip.NewWriter(dst)
	defer w.Close()

	w.Write(fileBytes)

	return nil
}

// Unzips file from uploads directory and return the uncompressed bytes
func (app *application) unzip(file string) ([]byte, error) {
	zipFile, err := os.Open(fmt.Sprintf("uploads/%s.gz", file))
	if err != nil {
		return nil, errors.New("unable to open zip file")
	}
	defer zipFile.Close()

	gz, err := gzip.NewReader(zipFile)
	if err != nil {
		return nil, errors.New("unable to create gzip reader from zip file")
	}
	defer gz.Close()

	bytes, err := io.ReadAll(gz)
	return bytes, err
}

func (app *application) tryHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		fmt.Println("Error while parsing multipart form")
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

		err = app.zip(header.Filename, file)
		if err != nil {
			app.logger.Printf("Unable to zip file: %s, %v\n", err.Error(), err)
			return
		}

		fmt.Println("Uploaded file")
	}
}

func (app *application) tryUploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/target.html")
	tmpl.Execute(w, nil)
}

func (app *application) ConvertToIST(inTime time.Time) time.Time {
	istTime := time.Hour*5 + time.Minute*30
	outTime := inTime.Local().Add(-istTime)
	return outTime
}
