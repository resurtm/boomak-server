package handlers

import (
	"path/filepath"
	"github.com/resurtm/boomak-server/common"
	"github.com/resurtm/boomak-server/user"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
	"strings"
)

func processHandlerData(container interface{}, schema string, w http.ResponseWriter, r *http.Request) bool {
	// step 1 - receive data
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to parse incoming data"))
		log.WithFields(log.Fields{
			"schema": schema,
			"err":    err,
		}).Warn("unable to parse incoming data")
		return false
	}

	// step 2 - validate schema
	schemaPath := "file://" + filepath.Join(common.CurrentDir(), jsonSchemaDefaultDirectory, schema+".json")
	result, err := gojsonschema.Validate(gojsonschema.NewReferenceLoader(schemaPath), gojsonschema.NewGoLoader(data))
	if err != nil || !result.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable validate data"))
		log.WithFields(log.Fields{
			"schema": schema,
			"data":   data,
			"err":    err,
		}).Warn("unable to validate json schema")
		return false
	}

	// step 3 - decode to the necessary structure
	if err := mapstructure.Decode(data, container); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot process your data"))
		log.WithFields(log.Fields{
			"schema": schema,
			"data":   data,
			"err":    err,
		}).Warn("unable to decode with mapstructure")
		return false
	}

	return true
}

func findUserByRequest(w http.ResponseWriter, r *http.Request) *user.User {
	data, ok := r.Header["Authorization"]
	if !ok || len(data) != 1 || !strings.Contains(data[0], "bearer ") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("valid Authorization header has not been set"))
		log.Warn("valid Authorization header has not been set")
		return nil
	}

	parts := strings.Split(data[0], " ")
	if len(parts) != 2 {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("valid Authorization header has not been set"))
		log.Warn("valid Authorization header has not been set")
		return nil
	}

	claims, err := user.CheckJWT(parts[1])
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("check of Authorization header has failed"))
		log.Warn("check of Authorization header has failed")
		return nil
	}

	username, email := claims["username"].(string), claims["email"].(string)
	if usr, err := user.FindByUsernameAndEmail(username, email, nil); err != nil || usr == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("user by given email and/or username cannot be found"))
		log.WithFields(log.Fields{
			"err": err,
			"username": username,
			"email": email,
		}).Warn("user by given email and/or username cannot be found")
		return nil
	} else {
		return usr
	}
}
