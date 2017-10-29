package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/auth", http.HandlerFunc(authHandler)).Methods("POST")
	r.Handle("/register", http.HandlerFunc(signupHandler)).Methods("POST")
	r.Handle("/validate", http.HandlerFunc(validateHandler)).Methods("POST")

	return r
}
