package database

import (
	"gopkg.in/mgo.v2"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/tools"
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
	Id                     int    `json:"id" bson:"_id,omitempty"`
	Username               string `json:"username"`
	Password               string `json:"password"`
	Email                  string `json:"email"`
	EmailVerified          bool   `json:"email_verified" bson:"email_verified"`
	EmailVerificationToken string `json:"email_verification_token" bson:"email_verification_token"`
}

func (user *User) MakeEmailNonVerified() {
	token, err := tools.GenerateRandomString(int(config.Config().Mailing.VerificationTokenLength))
	if err != nil {
		panic(err)
	}

	user.EmailVerified = false
	user.EmailVerificationToken = token
}
