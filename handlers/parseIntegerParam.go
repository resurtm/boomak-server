package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func parseIntegerParam(paramName string, w http.ResponseWriter, r *http.Request) (int, bool) {
	data, ok := r.URL.Query()[paramName]
	if !ok || len(data) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(paramName + " parameter has not been set"))
		log.Warn(paramName + " parameter has not been set")
		return 0, false
	}
	if value, err := strconv.Atoi(data[0]); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(paramName + " parameter must be non negative integer"))
		log.Warn(paramName + " parameter must be non negative integer")
		return 0, false
	} else {
		return value, true
	}
}
