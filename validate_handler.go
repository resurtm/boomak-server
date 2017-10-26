package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	var authToken string
	if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		authToken = string(bodyBytes)
	}

	// check auth token, step 1
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(Config.Security.JwtSigningKey), nil
	})

	// check auth token, step 2
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// check user in db
	query := bson.M{"username": claims["username"], "email": claims["email"]}
	if count, err := UserCol.Find(query).Count(); count != 1 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
