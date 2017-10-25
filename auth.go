package main

import (
	"time"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"fmt"
)

func authHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	//claims["admin"] = true
	//claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(config.Security.JwtSigningKey))
	fmt.Println(config.Security.JwtSigningKey)
	fmt.Println(tokenString)
	fmt.Println(err)

	w.Write([]byte(tokenString))
}
