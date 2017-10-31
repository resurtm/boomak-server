package handlers

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/middleware"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.Handle("/v1/login", http.HandlerFunc(loginHandler)).Methods("POST")
	r.Handle("/v1/register", http.HandlerFunc(registerHandler)).Methods("POST")
	r.Handle("/v1/check", http.HandlerFunc(checkHandler)).Methods("POST")
	r.Handle("/v1/get-settings", middleware.Auth(http.HandlerFunc(getSettingsHandler))).Methods("GET")
	r.Handle("/v1/verify-email", middleware.Auth(http.HandlerFunc(verifyEmailHandler))).Methods("POST")

	r.Handle("/v1/test-action", middleware.Auth(http.HandlerFunc(testActionHandler))).Methods("POST")
	if cfg.C().Mailing.EnableTestMailer {
		r.Handle("/v1/test-email", http.HandlerFunc(testEmailHandler)).Methods("POST")
	}

	h := setupCORS(r)
	if cfg.C().Server.Debug {
		h = handlers.LoggingHandler(os.Stdout, h)
	}
	return h
}

func setupCORS(handler http.Handler) http.Handler {
	s := cfg.C().CORS
	c := cors.New(cors.Options{
		AllowedOrigins: s.Origins,
		AllowedHeaders: s.Headers,
		Debug:          s.Debug,
	})
	return c.Handler(handler)
}
