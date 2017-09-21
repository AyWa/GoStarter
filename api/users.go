package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aywa/goNotify/auth"
	"github.com/aywa/goNotify/db"
	"github.com/gorilla/mux"
)

// protected
func getUserProfile(c auth.Claims) ErrorHandler {
	return func(w http.ResponseWriter, r *http.Request) (httpE HTTPError) {
		u, err := db.GetPrivateUser(c.Email)
		if err != nil {
			return HTTPError{err, "user not found", 404}
		}
		json.NewEncoder(w).Encode(u)
		return httpE
	}
}

// unprotected route
func postNewUser(w http.ResponseWriter, r *http.Request) (httpE HTTPError) {
	var u db.User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := d.Decode(&u)
	if err != nil {
		return HTTPError{err, "couldn't read the information", 403}
	}
	err = db.CreateUser(&u)
	if err != nil {
		return HTTPError{err, "Couldn't add you to db", 403}
	}
	return httpE
}

func login(w http.ResponseWriter, r *http.Request) (httpE HTTPError) {
	var ul db.User
	d := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := d.Decode(&ul)
	if err != nil {
		return HTTPError{err, "couldn't read the information", 403}
	}
	ul, err = db.CheckCredential(ul)
	if err != nil {
		return HTTPError{err, "email or password are wrong", 403}
	}
	token, err := auth.GetToken(ul.Email, ul.FirstName, time.Hour*24)
	if err != nil {
		return HTTPError{err, "server error please try again", 403}
	}
	json.NewEncoder(w).Encode(token)
	return httpE
}

// Users return all the route related to the users
func users(myRouter *mux.Router) {
	// unprotected route
	myRouter.Path("/users").Methods("POST").Handler(ErrorHandler(postNewUser))
	myRouter.Path("/users/login").Methods("POST").Handler(ErrorHandler(login))
	// protected route
	myRouter.Path("/users/profile").Methods("GET").Handler(AuthHandler(getUserProfile))
}
