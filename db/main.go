package db

import (
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/config"
)

type Session struct {
	*mgo.Session
}

var mainSession *Session

// get rid of DB(...) call when trying to achieve *mgo.Collection instance
func (session *Session) C(name string) *mgo.Collection {
	return session.DB(config.C().Database.Name).C(name)
}

func New() *Session {
	if mainSession == nil {
		if newMainSession, err := mgo.Dial(config.C().DSN()); err != nil {
			panic(err)
		} else {
			mainSession = &Session{Session: newMainSession}
		}
	}
	return &Session{Session: mainSession.New()}
}
