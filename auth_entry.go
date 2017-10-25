package main

import (
	"github.com/xeipuuv/gojsonschema"
	"path/filepath"
	"os"
	"log"
)

type AuthEntry struct {
	Username string
	Password string
}

func ValidateAuthEntryJson(json map[string]interface{}) []gojsonschema.ResultError {
	currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath.Join(currDir, "auth_entry_schema.json"))
	documentLoader := gojsonschema.NewGoLoader(json)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}

	if result.Valid() {
		return nil
	} else {
		return result.Errors()
	}
}
