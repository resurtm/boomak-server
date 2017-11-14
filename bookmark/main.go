package bookmark

import (
	"gopkg.in/mgo.v2/bson"
)

type Bookmark struct {
	Id   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User bson.ObjectId `json:"user" bson:"user"`
	Url  string        `json:"url" bson:"url"`
}
