package handlers

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/database"
)

func validateHandler(w http.ResponseWriter, r *http.Request) {
	var authToken string
	if bytes, err := ioutil.ReadAll(r.Body); len(bytes) == 0 || err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		authToken = string(bytes)
	}

	// check auth token, step 1
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config().Security.JWTSigningKey), nil
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check auth token, step 2
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session := database.New()
	defer session.Close()

	query := bson.M{"username": claims["username"], "email": claims["email"]}
	if count, err := session.Col("user").Find(query).Count(); count != 1 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
