package main

import (
	"encoding/json"
	"lp3/internal/data"
	"net/http"
	"strconv"
)

// Writes JSON data to client
func (app *application) writeJSON(w http.ResponseWriter, data interface{}) error {
	jsonRes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)

	return nil
}

func (app *application) getUserFromCookie(r *http.Request) (*data.User, error) {
	cookie, err := r.Cookie("user-cookie")
	if err != nil {
		app.logger.Printf("Unable to retrieve user ID from cookie: %v\n", err)
		return nil, err
	}
	userId, err := strconv.Atoi(cookie.Value)
	if err != nil {
		app.logger.Printf("Unable to convert cookie Value to Integer: %v\n", err)
		return nil, err
	}
	user, err := app.model.Um.GetUserById(userId)
	if err != nil {
		app.logger.Printf("Unable to retrieve user from User ID: %v\n", err)
		return nil, err
	}

	return user, nil
}
