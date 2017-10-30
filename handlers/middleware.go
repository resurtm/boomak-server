package handlers

import (
	"net/http"
	s "strings"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/database"
	"gopkg.in/mgo.v2/bson"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check header, part 1
		data, ok := r.Header["Authorization"]
		if !ok || len(data) != 1 || !s.Contains(data[0], "bearer ") {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// check header, part 2
		parts := s.Split(data[0], " ")
		if len(parts) != 2 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// validate token
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Config().Security.JWTSigningKey), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// check user
		session := database.New()
		defer session.Close()

		query := bson.M{"$and": []bson.M{
			{"username": claims["username"]},
			{"email": claims["email"]},
		}}
		if n, err := session.Col("user").Find(query).Count(); n == 0 || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
