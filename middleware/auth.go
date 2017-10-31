package middleware

import (
	"net/http"
	s "strings"
	"github.com/resurtm/boomak-server/user"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, ok := r.Header["Authorization"]
		if !ok || len(data) != 1 || !s.Contains(data[0], "bearer ") {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		parts := s.Split(data[0], " ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims, err := user.CheckJWT(parts[1])
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		username, email := claims["username"].(string), claims["email"].(string)
		if exists, err := user.ExistsByUsernameAndEmail(username, email, nil); err != nil || !exists {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
