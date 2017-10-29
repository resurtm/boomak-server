package database

import (
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/config"
)

type sessionType struct {
	S *mgo.Session
}

func (session *sessionType) C(name string) *mgo.Collection {
	return session.S.DB(config.Config().Database.Name).C(name)
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
