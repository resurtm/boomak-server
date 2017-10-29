package handlers

import (
	"net/http"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("validate handler"))
}
