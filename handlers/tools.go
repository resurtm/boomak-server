package handlers

import (
	"path/filepath"
	"net/http"
	"github.com/xeipuuv/gojsonschema"
	"github.com/resurtm/boomak-server/common"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func processHandlerData(container interface{}, schema string, w http.ResponseWriter, r *http.Request) bool {
	// step 1 - receive data
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	// step 2 - validate schema
	schemaPath := "file://" + filepath.Join(common.CurrentDir(), common.JSONSchemaDir, schema+".json")
	result, err := gojsonschema.Validate(gojsonschema.NewReferenceLoader(schemaPath), gojsonschema.NewGoLoader(data))
	if err != nil || !result.Valid() {
		w.WriteHeader(http.StatusBadRequest)
		return false
	}

	// step 3 - decode to the necessary structure
	if err := mapstructure.Decode(data, container); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}

	return true
}
