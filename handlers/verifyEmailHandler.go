package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/tools"
	"io/ioutil"
)

func verifyEmailHandler(w http.ResponseWriter, r *http.Request) {
	u := tools.FindUserByResponseWriter(w, r)
	if u == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil || len(bytes) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := u.VerifyEmail(string(bytes), nil); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
