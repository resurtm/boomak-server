package main

import (
	"net/http"

	"encoding/json"
	"fmt"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var document map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&document); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	res, _ := ValidateUserJson(document)
	fmt.Printf("\n%+v\n", res)
	fmt.Printf("\n%+v\n", document)

	if res {
		w.Write([]byte("good"))
	} else {
		w.Write([]byte("bad"))
	}
}
