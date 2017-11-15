package handlers

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/resurtm/boomak-server/bookmark"
	"github.com/resurtm/boomak-server/db"
	"encoding/json"
)

// todo: fixme: make pagination more efficient
// https://github.com/icza/minquery
// https://github.com/icza/minquery/pull/1
// https://stackoverflow.com/questions/40796666/need-to-use-pagination-in-mgo
// https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	session := db.New()
	defer session.Close()

	usr := findUserByRequest(w, r, session)
	if usr == nil {
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

	bookmarks, err := bookmark.FindByUserID(usr.Id.Hex(), offset, limit, session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch bookmarks"))
		log.WithFields(log.Fields{"user": usr, "offset": offset, "limit": limit, "err": err}).Warn("unable to fetch bookmarks")
		return
	}

	bookmarksCount, err := bookmark.CountByUserId(usr.Id.Hex(), session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to fetch bookmarks count"))
		log.WithFields(log.Fields{"user": usr, "offset": offset, "limit": limit, "err": err}).Warn("unable to fetch bookmarks count")
		return
	}

	respData := struct {
		Count     int                 `json:"count"`
		Bookmarks []bookmark.Bookmark `json:"bookmarks"`
	}{bookmarksCount, bookmarks}

	if resp, err := json.Marshal(respData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to prepare response data"))
		log.WithFields(log.Fields{"user": usr, "offset": offset, "limit": limit, "err": err}).Warn("unable to prepare response data")
	} else {
		w.Write(resp)
	}
}
