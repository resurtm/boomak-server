package handlers

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func getSettingsHandler(w http.ResponseWriter, r *http.Request) {
	usr := findUserByRequest(w, r)
	if usr == nil {
		return
	}

	data := struct {
		Username      string `json:"username"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	}{
		Username:      usr.Username,
		Email:         usr.Email,
		EmailVerified: usr.EmailVerified,
	}

	if resp, err := json.Marshal(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to prepare response data"))
		log.WithFields(log.Fields{
			"data": data,
			"user": usr,
		}).Warn("unable to marshal response data")
	} else {
		w.Write(resp)
	}
}
