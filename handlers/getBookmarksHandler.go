package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/resurtm/boomak-server/bookmark"
	"encoding/json"
)

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.URL.Query()["user_id"]
	if !ok || len(userID) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user ID parameter has not been set"))
		log.Warn("user ID parameter has not been set")
		return
	}

	offset, ok := parseIntegerParam("offset", w, r)
	if !ok {
		return
	}
	limit, ok := parseIntegerParam("limit", w, r)
	if !ok {
		return
	}

	bookmarks, err := bookmark.FindByUserID(userID[0], offset, limit, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch bookmarks"))
		log.WithFields(log.Fields{
			"user_id": userID,
			"offset":  offset,
			"limit":   limit,
			"err":     err,
		}).Warn("unable to fetch bookmarks")
		return
	}

	if resp, err := json.Marshal(bookmarks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to prepare response data"))
		log.WithFields(log.Fields{
			"user_id": userID,
			"offset":  offset,
			"limit":   limit,
			"err":     err,
		}).Warn("unable to prepare response data")
	} else {
		w.Write(resp)
	}
}
