package user

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/resurtm/boomak-server/common"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/db"
)

func FindByUsername(username string, session *db.Session) (*User, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	var user User
	if err := session.C(common.UserCollectionName).Find(bson.M{"username": username}).One(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func ExistsByUsernameAndEmail(username string, email string, session *db.Session) (bool, error) {
	return existsByUsernameAndOrEmail(username, email, session, "$and")
}

func ExistsByUsernameOrEmail(username string, email string, session *db.Session) (bool, error) {
	return existsByUsernameAndOrEmail(username, email, session, "$or")
}

func existsByUsernameAndOrEmail(username string, email string, session *db.Session, compOp string) (bool, error) {
	if session == nil {
		session = db.New()
		defer session.Close()
	}
	query := bson.M{compOp: []bson.M{
		{"username": username},
		{"email": email},
	}}
	if n, err := session.C(common.UserCollectionName).Find(query).Count(); err != nil {
		return false, err
	} else {
		return n != 0, nil
	}
}

func CheckJWT(authToken string) (jwt.MapClaims, error) {
	// check auth token, step 1
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.C().Security.JWTSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	// check auth token, step 2
	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return nil, fmt.Errorf("cannot validate JWT token")
	} else {
		return claims, nil
	}
}
