package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/user"
	log "github.com/sirupsen/logrus"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	token, ok := r.URL.Query()["token"]
	if !ok || len(token) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("auth JWT token has not been provided"))
		log.Warn("auth JWT token has not been provided")
		return
	}

	claims, err := user.CheckJWT(token[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid JWT token has been provided"))
		log.WithFields(log.Fields{"err": err}).Warn("invalid JWT token has been provided")
		return
	}

	username, email := claims["username"].(string), claims["email"].(string)
	if exists, err := user.ExistsByUsernameAndEmail(username, email, nil); err != nil || !exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid JWT token has been provided"))
		log.WithFields(log.Fields{
			"err":      err,
			"username": username,
			"email":    email,
		}).Warn("invalid JWT token has been provided")
	}
}
