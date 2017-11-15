package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/bookmark"
	"github.com/resurtm/boomak-server/db"
	log "github.com/sirupsen/logrus"
)

func setBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var bookm bookmark.Bookmark
	if !processHandlerData(&bookm, "bookmark", w, r) {
		return
	}

	session := db.New()
	defer session.Close()

	usr := findUserByRequest(w, r, session)
	if usr == nil {
		return
	}

	bookm.UserId = usr.Id
	if err := bookm.Create(session); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to create a new bookmark"))
		log.WithFields(log.Fields{"err": err, "bookmark": bookm}).Warn("unable to create a new bookmark")
		return
	}
}
