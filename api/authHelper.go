package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/aywa/goNotify/auth"
)

// HTTPError is a custum error struct containing: E (error),
// M: http message (string), HTTPCode (int)
type HTTPError struct {
	E        error
	M        string
	HTTPCode int
}

// ErrorHandler is a type which implement the ServeHTTP methods requirement
// for the http.Handler interface
type ErrorHandler func(w http.ResponseWriter, req *http.Request) HTTPError

func (h ErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// err handler
	httpErr := h(w, r)
	if httpErr.E != nil {
		log.Output(1, httpErr.E.Error()+" HTTPCode: "+strconv.Itoa(httpErr.HTTPCode))
		http.Error(w, httpErr.M, httpErr.HTTPCode)
		return
	}
}

// AuthHandler is a type which take a function which return a ErrorHandler type
// AuthHandler check if the user is Authorize before running the ErrorHandler
type AuthHandler func(c auth.Claims) ErrorHandler

func (a AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	claims, err := auth.GetUserFromAuth(r.Header.Get("Authorization"))
	if err != nil {
		log.Output(1, err.Error())
		w.WriteHeader(401)
		return
	}
	a(claims).ServeHTTP(w, r)
}
