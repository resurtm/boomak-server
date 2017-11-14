package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/resurtm/boomak-server/config"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"path/filepath"
	"github.com/resurtm/boomak-server/common"
	"github.com/xeipuuv/gojsonschema"
	"github.com/mitchellh/mapstructure"
)

const jsonSchemaDefaultDirectory = "jsonSchema"

func New() http.Handler {
	log.Info("creating handlers object")
	r := mux.NewRouter()

	r.Handle("/v1/login", http.HandlerFunc(loginHandler)).Methods("POST")
	/*r.Handle("/v1/check", http.HandlerFunc(checkHandler)).Methods("GET")
	r.Handle("/v1/register", http.HandlerFunc(registerHandler)).Methods("POST")

	r.Handle("/v1/get-settings", middleware.Auth(http.HandlerFunc(getSettingsHandler))).Methods("GET")
	r.Handle("/v1/verify-email", middleware.Auth(http.HandlerFunc(verifyEmailHandler))).Methods("POST")*/

	r.Handle("/v1/test-guest", http.HandlerFunc(testGuestHandler)).Methods("POST")
	/*r.Handle("/v1/test-auth", middleware.Auth(http.HandlerFunc(testActionHandler))).Methods("POST")
	if cfg.C().Mailing.EnableTestMailer {
		r.Handle("/v1/test-email", http.HandlerFunc(testEmailHandler)).Methods("POST")
	}*/

	c := cors.New(cors.Options{
		AllowedOrigins: config.C().CORS.Origins,
		AllowedHeaders: config.C().CORS.Headers,
	})
	return handlers.LoggingHandler(log.StandardLogger().Writer(), c.Handler(r))
}

func processHandlerData(container interface{}, schema string, w http.ResponseWriter, r *http.Request) bool {
	// step 1 - receive data
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("unable to parse incoming data"))
		log.WithFields(log.Fields{
			"schema": schema,
			"err": err,
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
			"data": data,
			"err": err,
		}).Warn("unable to validate json schema")
		return false
	}

	// step 3 - decode to the necessary structure
	if err := mapstructure.Decode(data, container); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("cannot process your data"))
		log.WithFields(log.Fields{
			"schema": schema,
			"data": data,
			"err": err,
		}).Warn("unable to decode with mapstructure")
		return false
	}

	return true
}
