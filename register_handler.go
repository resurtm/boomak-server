package main

import (
	"net/http"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
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

	find := bson.M{"$or": []bson.M{
		{"username": user.Username},
		{"email": user.Email},
	}}
	if n, err := UserCol.Find(find).Count(); n != 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

	SignupMails <- user
}
