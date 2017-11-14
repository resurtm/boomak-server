package bookmark

import (
	"github.com/resurtm/boomak-server/db"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/resurtm/boomak-server/user"
	"errors"
)

func FindByUserID(userID string, offset int, limit int, session *db.Session) ([]Bookmark, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}

	exists, err := user.ExistsByID(userID, session)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("non existing user ID has been provided")
	}

	bookmarks := []Bookmark{}
	query := bson.M{"user": bson.ObjectIdHex(userID)}
	err = session.C("bookmark").Find(query).Sort("_id").Skip(offset).Limit(limit).All(&bookmarks)
	if err != nil {
		return nil, err
	} else {
		return bookmarks, nil
	}
}

func GenerateBookmarks(bookmarkCount uint, userID string, session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}

	exists, err := user.ExistsByID(userID, session)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("non existing user ID has been provided")
	}

	for i := uint(0); i < bookmarkCount; i++ {
		bookmark := Bookmark{
			Id:   bson.NewObjectId(),
			User: bson.ObjectIdHex(userID),
			Url:  fmt.Sprintf("http://%s/", fake.DomainName()),
		}

		if err := session.C("bookmark").Insert(bookmark); err != nil {
			return err
		}
	}

	return nil
}
