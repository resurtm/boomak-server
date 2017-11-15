package handlers

import (
	"net/http"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if findUserByRequest(w, r, nil) == nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
