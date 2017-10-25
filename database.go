package main

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

var DbSession *mgo.Session
var Db *mgo.Database
var UserCol *mgo.Collection

func ConnectToDb() *mgo.Session {
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/%s",
		Config.Database.User,
		Config.Database.Password,
		Config.Database.Hostname,
		Config.Database.Port,
		Config.Database.Name,
	)

	session, err := mgo.Dial(dsn)
	if err != nil {
		panic(err)
	}

	DbSession = session
	DbSession.SetMode(mgo.Monotonic, true)
	Db = DbSession.DB(Config.Database.Name)
	UserCol = Db.C("user")

	return DbSession
}
