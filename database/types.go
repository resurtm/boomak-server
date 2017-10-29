package database

import (
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/config"
)

type sessionType struct {
	*mgo.Session
}

// Get rid of DB(...) call when trying to achieve *mgo.Collection instance.
func (session *sessionType) Col(name string) *mgo.Collection {
	return session.DB(config.Config().Database.Name).C(name)
}

type AuthEntry struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
