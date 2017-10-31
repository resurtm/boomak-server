package user

import (
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/common"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/resurtm/boomak-server/mailing/jobs"
	"github.com/resurtm/boomak-server/mailing/types"
	"github.com/resurtm/boomak-server/db"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username"`
	Password string        `json:"password"`

	Email                  string `json:"email"`
	EmailVerified          bool   `json:"email_verified" bson:"email_verified"`
	EmailVerificationToken string `json:"email_verification_token" bson:"email_verification_token"`
}

func (user *User) MakeEmailNonVerified(commit bool, sendEmail bool, session *db.Session) error {
	token, err := common.GenerateRandomString(int(cfg.C().Mailing.VerificationTokenLength))
	if err != nil {
		return err
	}
	user.EmailVerified = false
	user.EmailVerificationToken = token
	if commit {
		if session == nil {
			session = db.New()
			defer session.Close()
		}
		query := bson.M{"$and": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		}}
		change := bson.M{"$set": bson.M{
			"email_verified":           user.EmailVerified,
			"email_verification_token": user.EmailVerificationToken,
		}}
		if err := session.C("user").Update(query, change); err != nil {
			return err
		}
	}
	if sendEmail {
		jobs.MailJobsQueue <- types.MailJob{Kind: types.EmailVerifyMailJob, Payload: *user}
	}
	return nil
}

func (user *User) SetRawPassword(password string) bool {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return false
	} else {
		user.Password = string(hashed)
		return true
	}
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(cfg.C().Security.JWTSigningKey))
}

func (user *User) Update(session *db.Session) error {
	return nil
}

func (user *User) Create(session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	user.Id = bson.NewObjectId()
	return session.C("user").Insert(user)
}
