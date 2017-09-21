package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var myRouter *mux.Router

// StartAPI start the api server
func StartAPI(port string) {
	myRouter = mux.NewRouter().Path("/api").Subrouter()
	users(myRouter)
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}
