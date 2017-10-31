package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/tools"
	"encoding/json"
)

func getSettingsHandler(w http.ResponseWriter, r *http.Request) {
	u := tools.FindUserByResponseWriter(w, r)
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := struct {
		Username      string `json:"username"`
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified" bson:"email_verified"`
	}{
		Username:      u.Username,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
	}

	if response, err := json.Marshal(data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(response)
	}
}
