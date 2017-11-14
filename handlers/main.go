package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/resurtm/boomak-server/config"
	log "github.com/sirupsen/logrus"
)

const jsonSchemaDefaultDirectory = "jsonSchema"

func New() http.Handler {
	log.Info("creating handlers object")
	r := mux.NewRouter()

	r.Handle("/v1/login", http.HandlerFunc(loginHandler)).Methods("POST")
	r.Handle("/v1/check", http.HandlerFunc(checkHandler)).Methods("GET")
	r.Handle("/v1/register", http.HandlerFunc(registerHandler)).Methods("POST")

	/*r.Handle("/v1/get-settings", middleware.Auth(http.HandlerFunc(getSettingsHandler))).Methods("GET")
	r.Handle("/v1/verify-email", middleware.Auth(http.HandlerFunc(verifyEmailHandler))).Methods("POST")*/

	r.Handle("/v1/test-guest", http.HandlerFunc(testActionHandler)).Methods("POST")
	r.Handle("/v1/test-auth", authMiddleware(http.HandlerFunc(testActionHandler))).Methods("POST")
	/*if cfg.C().Mailing.EnableTestMailer {
		r.Handle("/v1/test-email", http.HandlerFunc(testEmailHandler)).Methods("POST")
	}*/

	c := cors.New(cors.Options{
		AllowedOrigins: config.C().CORS.Origins,
		AllowedHeaders: config.C().CORS.Headers,
	})
	return handlers.LoggingHandler(log.StandardLogger().Writer(), c.Handler(r))
}
