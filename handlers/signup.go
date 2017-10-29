package handlers

import (
	"net/http"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup handler"))
}
