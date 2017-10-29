package handlers

import (
	"net/http"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth handler"))
}
