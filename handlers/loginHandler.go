package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/types"
	"github.com/resurtm/boomak-server/user"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var authEntry types.AuthEntry
	if !processHandlerData(&authEntry, "authEntry", w, r) {
		return
	}

	u, err := user.FindByUsername(authEntry.Username, nil)
	if err != nil || !u.CheckPassword(authEntry.Password) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if token, err := u.GenerateJWT(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(token))
	}
}
