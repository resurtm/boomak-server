package handlers

import (
	"net/http"
	"io/ioutil"
	log "github.com/sirupsen/logrus"
)

func testActionHandler(w http.ResponseWriter, r *http.Request) {
	if bytes, err := ioutil.ReadAll(r.Body); err != nil || len(bytes) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to read incoming data"))
		log.WithFields(log.Fields{"err": err}).Warn("unable to parse incoming data")
	} else {
		w.Write(bytes)
	}
}
