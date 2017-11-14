package user

import (
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/common"
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/resurtm/boomak-server/db"
	"errors"
	log "github.com/sirupsen/logrus"
	mailing "github.com/resurtm/boomak-server/mailing/base"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`

	Email                  string `json:"email" bson:"email"`
	EmailVerified          bool   `json:"email_verified" bson:"email_verified"`
	EmailVerificationToken string `json:"email_verification_token" bson:"email_verification_token"`
}

func (user *User) MakeEmailNonVerified(commit bool, sendEmail bool, session *db.Session) error {
	token, err := common.GenerateRandomString(int(config.C().Mailing.VerificationTokenLength))
	if err != nil {
		return err
	}
	user.EmailVerified = false
	user.EmailVerificationToken = token[1:len(token)-3]

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
		mailing.EnqueueEmailVerifyMailJob(user)
	}
	return nil
}

func (user *User) SetRawPassword(password string) error {
	if hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		return err
	} else {
		user.Password = string(hashed)
		return nil
	}
}

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func (user *User) GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(config.C().Security.JWTSigningKey))
}

func (user *User) Create(session *db.Session) error {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	user.Id = bson.NewObjectId()
	return session.C("user").Insert(user)
}

func (user *User) VerifyEmail(key string, session *db.Session) error {
	log.WithFields(log.Fields{
		"valid_token": user.EmailVerificationToken,
		"checked_token": key,
	}).Info("checking user's email verification token")
	if user.EmailVerificationToken != key {
		return errors.New("invalid token has been passed")
	}

	if session == nil {
		session = db.New()
		defer session.Close()
	}
	query := bson.M{"$and": []bson.M{
		{"username": user.Username},
		{"email": user.Email},
	}}
	change := bson.M{"$set": bson.M{
		"email_verified":           true,
		"email_verification_token": nil,
	}}
	if err := session.C("user").Update(query, change); err != nil {
		return err
	}

	mailing.EnqueueSignupFinishedMailJob(user)
	return nil
}
