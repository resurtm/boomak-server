package handlers

import (
	"net/http"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/mitchellh/mapstructure"
	"github.com/resurtm/boomak-server/database"
	"github.com/resurtm/boomak-server/mailing"
)

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validateHandlerData(data, "user", w) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user database.User
	if err := mapstructure.Decode(data, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	session := database.New()
	defer session.Close()

	query := bson.M{"$or": []bson.M{
		{"username": user.Username},
		{"email": user.Email},
	}}
	if n, err := session.Col("user").Find(query).Count(); n != 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		user.Password = string(hashed)
	}

	if err := session.Col("user").Insert(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		mailing.SendSignupEmail(user)
	}
}
