package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/db"
	"github.com/resurtm/boomak-server/user"
	log "github.com/sirupsen/logrus"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var usr user.User
	if !processHandlerData(&usr, "user", w, r) {
		return
	}

	session := db.New()
	defer session.Close()

	if exists, err := user.ExistsByUsernameOrEmail(usr.Username, usr.Email, session); err != nil || exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user with such email and/or password already exists"))
		log.WithFields(log.Fields{
			"err": err,
			"user": usr,
		}).Warn("user with such email and/or password already exists")
		return
	}

	if err := usr.SetRawPassword(usr.Password); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to hash the given password"))
		log.WithFields(log.Fields{
			"err": err,
			"user": usr,
		}).Warn("unable to hash the given password")
		return
	}

	if err := usr.Create(session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to create a new user"))
		log.WithFields(log.Fields{
			"err": err,
			"user": usr,
		}).Warn("unable to create a new user")
		return
	}

	if err := usr.MakeEmailNonVerified(true, true, session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to initialize the new user's email"))
		log.WithFields(log.Fields{
			"err": err,
			"user": usr,
		}).Warn("unable to initialize the new user's email")
	}
}
