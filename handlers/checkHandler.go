package handlers

import (
	"net/http"
	"io/ioutil"
	"github.com/resurtm/boomak-server/user"
)

func checkHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil || len(bytes) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	claims, err := user.CheckJWT(string(bytes))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username, email := claims["username"].(string), claims["email"].(string)
	if exists, err := user.ExistsByUsernameAndEmail(username, email, nil); err != nil || !exists {
		w.WriteHeader(http.StatusBadRequest)
	}
}
