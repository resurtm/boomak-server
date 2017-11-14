package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/resurtm/boomak-server/bookmark"
)

func generateBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	data, ok := r.URL.Query()["user_id"]
	if !ok || len(data) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user ID parameter has not been set"))
		log.Warn("user ID parameter has not been set")
		return
	}

	if err := bookmark.GenerateBookmarks(200, data[0], nil); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		log.WithField("user_id", data[0]).Warn(err)
	}
}
