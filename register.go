package main

import (
	"net/http"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	// decode data
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// validate json data
	if err := ValidateUserJson(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON schema validation failed"))
		return
	}

	// decode user struct
	var user User
	if err := mapstructure.Decode(data, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user.Password = string(hashedPassword)

	// save to db
	collection := db.C("user")
	if err := collection.Insert(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
