package handlers

import (
	"path/filepath"
	"github.com/resurtm/boomak-server/common"
	"github.com/xeipuuv/gojsonschema"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/mitchellh/mapstructure"
	"encoding/json"
)

const jsonSchemaDefaultDirectory = "jsonSchema"

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
