package handlers

import (
	"net/http"
	"io/ioutil"
)

func testActionHandler(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.Write(bytes)
	}
}
