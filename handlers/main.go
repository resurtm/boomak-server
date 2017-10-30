package handlers

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	cfg "github.com/resurtm/boomak-server/config"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.Handle("/v1/login", http.HandlerFunc(authHandler)).Methods("POST")
	r.Handle("/v1/register", http.HandlerFunc(signupHandler)).Methods("POST")
	r.Handle("/v1/check", http.HandlerFunc(validateHandler)).Methods("POST")

	r.Handle("/v1/test-action", authMiddleware(http.HandlerFunc(testActionHandler))).Methods("POST")

	if cfg.Config().Mailing.EnableTestMailing {
		r.Handle("/v1/test-email", http.HandlerFunc(testEmailHandler)).Methods("POST")
	}

	h := setupCORS(r)
	if cfg.Config().Server.DebugOutput {
		h = handlers.LoggingHandler(os.Stdout, h)
	}
	return h
}

func setupCORS(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: cfg.Config().CORS.Origins,
		Debug:          cfg.Config().CORS.Debug,
	})
	return c.Handler(handler)
}
