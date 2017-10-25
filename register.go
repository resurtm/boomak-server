package main

import (
	"net/http"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := ValidateUserJson(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("JSON schema validation failed"))
		return
	}

	var user User
	if err := mapstructure.Decode(data, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	collection := db.C("user")
	if err := collection.Insert(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
