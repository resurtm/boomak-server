package main

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	data := DecodeHandlerData(w, r)
	if data == nil {
		return
	}

	if !ValidateHandlerData(w, data, "auth_entry") {
		return
	}

	var authEntry AuthEntry
	if err := mapstructure.Decode(data, &authEntry); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var user User
	if err := UserCol.Find(bson.M{"username": authEntry.Username}).One(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authEntry.Password)); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	// issue JWT token
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(Config.Security.JwtSigningKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tokenString))
}
