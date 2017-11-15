package handlers

import (
	"github.com/resurtm/boomak-server/user"
	"net/http"
	log "github.com/sirupsen/logrus"
	"strings"
	"github.com/resurtm/boomak-server/db"
)

func findUserByRequest(w http.ResponseWriter, r *http.Request, session *db.Session) *user.User {
	data, ok := r.Header["Authorization"]
	if !ok || len(data) != 1 || !strings.Contains(data[0], "bearer ") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("valid Authorization header has not been set"))
		log.Warn("valid Authorization header has not been set")
		return nil
	}

	parts := strings.Split(data[0], " ")
	if len(parts) != 2 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("valid Authorization header has not been set"))
		log.Warn("valid Authorization header has not been set")
		return nil
	}

	claims, err := user.CheckJWT(parts[1])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("check of Authorization header has failed"))
		log.Warn("check of Authorization header has failed")
		return nil
	}

	username, email := claims["username"].(string), claims["email"].(string)
	if usr, err := user.FindByUsernameAndEmail(username, email, session); err != nil || usr == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user by given email and/or username cannot be found"))
		log.WithFields(log.Fields{
			"err":      err,
			"username": username,
			"email":    email,
		}).Warn("user by given email and/or username cannot be found")
		return nil
	} else {
		return usr
	}
}
