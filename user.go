package main

import (
	"os"
	"path/filepath"
	"log"

	"github.com/xeipuuv/gojsonschema"
)

type User struct {
	Id       int
	Username string
	Email    string
}

func ValidateUserJson(json map[string]interface{}) []gojsonschema.ResultError {
	currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + filepath.Join(currDir, "user_schema.json"))
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
