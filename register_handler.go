package main

import (
	"net/http"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := DecodeHandlerData(w, r)
	if data == nil {
		return
	}

	if !ValidateHandlerData(w, data, "user") {
		return
	}

	var user User
	if err := mapstructure.Decode(data, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	user.Password = string(hashedPassword)

	if err := UserCol.Insert(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
