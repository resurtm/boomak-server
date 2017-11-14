package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/common"
	"github.com/resurtm/boomak-server/user"
	log "github.com/sirupsen/logrus"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var authEntry common.AuthEntry
	if !processHandlerData(&authEntry, "authEntry", w, r) {
		return
	}

	usr, err := user.FindByUsername(authEntry.Username, nil)
	if err != nil || !usr.CheckPassword(authEntry.Password) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username and/or password have been provided"))
		log.WithFields(log.Fields{"err": err}).Warn("invalid username and/or password have been provided")
		return
	}

	if token, err := usr.GenerateJWT(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to generate JWT token"))
		log.WithFields(log.Fields{"err": err}).Warn("unable to generate JWT token")
	} else {
		w.Write([]byte(token))
	}
}
