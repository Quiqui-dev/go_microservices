package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type parameters struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (app *Config) Authenticate(w http.ResponseWriter, r *http.Request) {

	params := parameters{}

	err := app.readJSON(r, &params)

	if err != nil {
		log.Println(err, params.Email)
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate user against db
	log.Printf("Email: %s \n Password: %s", params.Email, params.Password)
	dbUser, err := app.DB.GetUserByEmail(r.Context(), params.Email)

	if err != nil {
		log.Println(err, params.Email)
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// convert user to something
	user := databaseUserToUser(dbUser)

	valid, err := user.PasswordMatches(params.Password)

	if err != nil || !valid {
		log.Println("invalid")
		app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.EmailAddress),
		Data:    user,
	}

	log.Println("Sending response back to broker")
	app.writeJSON(w, http.StatusAccepted, resp)
}
