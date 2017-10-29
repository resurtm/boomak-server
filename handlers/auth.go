package handlers

import (
	"net/http"
	"github.com/mitchellh/mapstructure"
	"github.com/resurtm/boomak-server/database"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validateHandlerData(data, "authEntry", w) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var authEntry database.AuthEntry
	if err := mapstructure.Decode(data, &authEntry); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session := database.New()
	defer session.S.Close()

	var user database.User
	if err := session.C("user").Find(bson.M{"username": authEntry.Username}).One(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authEntry.Password)); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	if token, err := generateJWT(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(token))
	}
}
