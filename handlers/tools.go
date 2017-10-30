package handlers

import (
	"path/filepath"
	"net/http"
	"github.com/xeipuuv/gojsonschema"
	"github.com/resurtm/boomak-server/config"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/resurtm/boomak-server/database"
	"github.com/resurtm/boomak-server/tools"
)

func validateHandlerData(data map[string]interface{}, schema string, w http.ResponseWriter) bool {
	schemaPath := "file://" + filepath.Join(tools.CurrentDir(), config.Config().Security.JSONSchemaDir, schema+".json")
	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
	documentLoader := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}
	return result.Valid()
}

func generateJWT(user database.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString([]byte(config.Config().Security.JWTSigningKey))
}
