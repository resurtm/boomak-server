package bookmark

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/resurtm/boomak-server/db"
)

type Bookmark struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserId bson.ObjectId `json:"user" bson:"user,omitempty"`
	Url    string        `json:"url" bson:"url"`
}

func (bookmark *Bookmark) Create(session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	bookmark.Id = bson.NewObjectId()
	return session.C("bookmark").Insert(bookmark)
}

func (bookmark *Bookmark) Delete(session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	query := bson.M{"_id": bookmark.Id}
	return session.C("bookmark").Remove(query)
}
