package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/db"
	"github.com/resurtm/boomak-server/bookmark"
	log "github.com/sirupsen/logrus"
)

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	bookmarkId, ok := parseStringParam("id", w, r)
	if !ok {
		return
	}

	session := db.New()
	defer session.Close()

	usr := findUserByRequest(w, r, session)
	if usr == nil {
		return
	}

	bookm, err := bookmark.FindOneById(bookmarkId, usr.Id.Hex(), session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch bookmark"))
		log.WithFields(log.Fields{
			"bookmark_id": bookmarkId,
			"user":        usr,
			"err":         err,
		}).Warn("unable to fetch bookmark")
		return
	}

	if err := bookm.Delete(session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to delete bookmark"))
		log.WithFields(log.Fields{
			"bookmark_id": bookmarkId,
			"bookmark":    bookm,
			"user":        usr,
			"err":         err,
		}).Warn("unable to delete bookmark")
	}
}
