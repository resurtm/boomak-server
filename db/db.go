package db

import (
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/cfg"
)

type Session struct {
	*mgo.Session
}

// get rid of DB(...) call when trying to achieve *mgo.Collection instance
func (session *Session) C(name string) *mgo.Collection {
	return session.DB(cfg.C().Database.Name).C(name)
}

var mainSession *Session

func New() *Session {
	if mainSession == nil {
		if newMainSession, err := mgo.Dial(cfg.C().DSN()); err != nil {
			panic(err)
		} else {
			mainSession = &Session{Session: newMainSession}
		}
	}
	return &Session{Session: mainSession.New()}
}
