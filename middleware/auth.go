package middleware

import (
	"net/http"
	"github.com/resurtm/boomak-server/tools"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tools.FindUserByResponseWriter(w, r) == nil {
			return
		}
		next.ServeHTTP(w, r)
	})
}
