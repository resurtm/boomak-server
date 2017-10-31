package tools

import (
	"net/http"
	"strings"
	"github.com/resurtm/boomak-server/user"
)

func FindUserByResponseWriter(w http.ResponseWriter, r *http.Request) *user.User {
	data, ok := r.Header["Authorization"]
	if !ok || len(data) != 1 || !strings.Contains(data[0], "bearer ") {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}
	parts := strings.Split(data[0], " ")
	if len(parts) != 2 {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	claims, err := user.CheckJWT(parts[1])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return nil
	}

	username, email := claims["username"].(string), claims["email"].(string)
	if u, err := user.FindByUsernameAndEmail(username, email, nil); err != nil || u == nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	} else {
		return u
	}
}
