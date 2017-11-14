package handlers

import (
	"net/http"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	usr := findUserByRequest(w, r)
	if usr == nil {
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil || len(bytes) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to parse incoming data"))
		log.WithFields(log.Fields{
			"err":  err,
			"user": usr,
		}).Warn("unable to parse incoming data")
		return
	}

	if err := usr.VerifyEmail(string(bytes), nil); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("email verification code is invalid"))
		log.WithFields(log.Fields{
			"err":  err,
			"user": usr,
			"code": string(bytes),
		}).Warn("email verification code is invalid")
	}
}
