package user

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/db"
	"errors"
)

func FindByUsername(username string, session *db.Session) (*User, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	var user User
	query := bson.M{"username": username}
	if err := session.C("user").Find(query).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByUsernameAndEmail(username string, email string, session *db.Session) (*User, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	var user User
	query := bson.M{"username": username, "email": email}
	if err := session.C("user").Find(query).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func ExistsByUsernameAndEmail(username string, email string, session *db.Session) (bool, error) {
	return existsByUsernameEmail(username, email, session, "$and")
}

func ExistsByUsernameOrEmail(username string, email string, session *db.Session) (bool, error) {
	return existsByUsernameEmail(username, email, session, "$or")
}

func existsByUsernameEmail(username string, email string, session *db.Session, compOp string) (bool, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	query := bson.M{compOp: []bson.M{
		{"username": username},
		{"email": email},
	}}
	if n, err := session.C("user").Find(query).Count(); err != nil {
		return false, err
	} else {
		return n != 0, nil
	}
}

func ExistsByID(userID string, session *db.Session) (bool, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	query := bson.M{"_id": bson.ObjectId(userID)}
	if n, err := session.C("user").Find(query).Count(); err != nil {
		return false, err
	} else {
		return n != 0, err
	}
}

func CheckJWT(authToken string) (jwt.MapClaims, error) {
	// check auth token, step 1
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.C().Security.JWTSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	// check auth token, step 2
	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, errors.New("cannot validate JWT token")
	} else {
		return claims, nil
	}
}
