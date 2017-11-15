package bookmark

import (
	"github.com/resurtm/boomak-server/db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/resurtm/boomak-server/user"
	"errors"
	log "github.com/sirupsen/logrus"
)

// todo: fixme: make pagination more efficient
// https://github.com/icza/minquery
// https://github.com/icza/minquery/pull/1
// https://stackoverflow.com/questions/40796666/need-to-use-pagination-in-mgo
// https://stackoverflow.com/questions/40634865/efficient-paging-in-mongodb-using-mgo
func FindByUserID(userId string, offset int, limit int, session *db.Session) ([]Bookmark, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}

	exists, err := user.ExistsByID(userId, session)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("non existing user ID has been provided")
	}

	bookmarks := []Bookmark{}
	query := bson.M{"user": bson.ObjectId(userId)}

	log.WithFields(log.Fields{
		"user_id": userId,
		"query":   query,
	}).Debug("trying to find a set of bookmarks by user ID")

	err = session.C("bookmark").Find(query).Sort("_id").Skip(offset).Limit(limit).All(&bookmarks)
	if err != nil {
		return nil, err
	} else {
		return bookmarks, nil
	}
}

func FindOneById(bookmarkId string, userId string, session *db.Session) (*Bookmark, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}

	var bookmark Bookmark
	query := bson.M{"_id": bson.ObjectIdHex(bookmarkId), "user": bson.ObjectIdHex(userId)}

	log.WithFields(log.Fields{
		"bookmark_id": bookmarkId,
		"user_id":     userId,
		"query":       query,
	}).Debug("trying to find a single bookmark by ID")

	if err := session.C("bookmark").Find(query).One(&bookmark); err != nil {
		return nil, err
	}
	return &bookmark, nil
}

func GenerateBookmarks(bookmarkCount uint, userId string, session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}

	exists, err := user.ExistsByID(userId, session)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("non existing user ID has been provided")
	}

	for i := uint(0); i < bookmarkCount; i++ {
		bookmark := Bookmark{
			Id:     bson.NewObjectId(),
			UserId: bson.ObjectId(userId),
			Url:    fmt.Sprintf("http://%s/", fake.DomainName()),
		}

		if err := session.C("bookmark").Insert(bookmark); err != nil {
			return err
		}
	}

	return nil
}
