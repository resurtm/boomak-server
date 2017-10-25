package main

import (
	"net/http"
	"encoding/json"
	"path/filepath"
	"os"
	"github.com/xeipuuv/gojsonschema"
)

const jsonSchemaDir = "json_schema"

func DecodeHandlerData(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return nil
	}
	return data
}

func ValidateHandlerData(w http.ResponseWriter, jsonData map[string]interface{}, jsonSchema string) bool {
	currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	schemaPath := "file://" + filepath.Join(currDir, jsonSchemaDir, jsonSchema+".json")
	schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)
	documentLoader := gojsonschema.NewGoLoader(jsonData)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}
	return result.Valid()
}
