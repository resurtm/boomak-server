package main

import (
	"time"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	// decode data
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// validate json data
	if err := ValidateAuthEntryJson(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON schema validation failed"))
		return
	}

	// decode auth entry struct
	var authEntry AuthEntry
	if err := mapstructure.Decode(data, &authEntry); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// find user from the db
	var user User
	if err := db.C("user").Find(bson.M{"username": authEntry.Username}).One(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// check the password
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

	tokenString, err := token.SignedString([]byte(config.Security.JwtSigningKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(tokenString))
}
