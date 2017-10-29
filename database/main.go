package database

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/config"
)

var mainSession *sessionType

func New() *sessionType {
	if mainSession == nil {
		if session, err := mgo.Dial(dsn()); err != nil {
			panic(err)
		} else {
			mainSession = &sessionType{S: session}
		}
	}
	return &sessionType{S: mainSession.S.New()}
}

func dsn() string {
	var dsn string
	cfg := config.Config().Database
	if cfg.NoPassword {
		dsn = fmt.Sprintf(
			"mongodb://%s:%d/%s",
			cfg.Hostname,
			cfg.Port,
			cfg.Name,
		)
	} else {
		// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		dsn = fmt.Sprintf(
			"mongodb://%s:%s@%s:%d/%s",
			cfg.User,
			cfg.Password,
			cfg.Hostname,
			cfg.Port,
			cfg.Name,
		)
	}
	return dsn
}
