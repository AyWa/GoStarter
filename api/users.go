package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getPeopleEndpoint(w http.ResponseWriter, req *http.Request) func() {
	return func() {
		json.NewEncoder(w).Encode("aa: heyyy")
	}
}

func postNewUser(w http.ResponseWriter, req *http.Request) {

}

// Users return all the route related to the users
func users(myRouter *mux.Router) {
	myRouter.Path("/users").Methods("GET").HandlerFunc(isAuthenticated(getPeopleEndpoint))
	myRouter.Path("/users").Methods("POST").HandlerFunc(postNewUser)
}
