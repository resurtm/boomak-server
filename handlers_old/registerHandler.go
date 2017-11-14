package handlers_old

import (
	"net/http"
	"github.com/resurtm/boomak-server/db"
	"github.com/resurtm/boomak-server/user"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var u user.User
	if !processHandlerData(&u, "user", w, r) {
		return
	}

	session := db.New()
	defer session.Close()

	if exists, err := user.ExistsByUsernameOrEmail(u.Username, u.Email, session); err != nil || exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !u.SetRawPassword(u.Password) || u.Create(session) != nil ||
		u.MakeEmailNonVerified(true, true, session) != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
