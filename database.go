package main

import (
	"gopkg.in/mgo.v2"
)

var dbSession *mgo.Session
var db *mgo.Database

func connectToDb() *mgo.Session {
	session, err := mgo.Dial(config.Database.Hostname)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	dbSession = session
	db = dbSession.DB(config.Database.Name)
	return session
}
