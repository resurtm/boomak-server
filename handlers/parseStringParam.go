package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func parseStringParam(paramName string, w http.ResponseWriter, r *http.Request) (string, bool) {
	if data, ok := r.URL.Query()[paramName]; !ok || len(data) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(paramName + " parameter has not been set"))
		log.Warn(paramName + " parameter has not been set")
		return "", false
	} else {
		return data[0], true
	}
}
