package api

import (
	"net/http"

	"github.com/aywa/goNotify/auth"
)

// we add a closure to be sure that a protected route and a normal route has different return value
func isAuthenticated(f func(w http.ResponseWriter, req *http.Request) func()) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		err := auth.GetUserFromAuth(req.Header.Get("Authorization"))

		if err != nil {
			w.WriteHeader(401)
			return
		}

		f(w, req)()
	}
}
