package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/resurtm/boomak-server/config"
	"github.com/sirupsen/logrus"
	"log"
)

func New() http.Handler {
	logrus.Info("creating handlers object")
	r := mux.NewRouter()

	r.Handle("/v1/login", http.HandlerFunc(loginHandler)).Methods("POST")
	r.Handle("/v1/check", http.HandlerFunc(checkHandler)).Methods("GET")
	r.Handle("/v1/register", http.HandlerFunc(registerHandler)).Methods("POST")

	r.Handle("/v1/settings", authMiddleware(http.HandlerFunc(getSettingsHandler))).Methods("GET")
	r.Handle("/v1/verify-email", authMiddleware(http.HandlerFunc(verifyEmailHandler))).Methods("POST")

	r.Handle("/v1/bookmark", authMiddleware(http.HandlerFunc(getBookmarksHandler))).Methods("GET")
	r.Handle("/v1/bookmark", authMiddleware(http.HandlerFunc(setBookmarkHandler))).Methods("POST")
	r.Handle("/v1/bookmark", authMiddleware(http.HandlerFunc(deleteBookmarkHandler))).Methods("DELETE")

	r.Handle("/v1/test-guest", http.HandlerFunc(testActionHandler)).Methods("POST")
	r.Handle("/v1/test-auth", authMiddleware(http.HandlerFunc(testActionHandler))).Methods("POST")

	if config.C().Security.EnableFaker {
		r.Handle("/v1/generate-bookmarks", http.HandlerFunc(generateBookmarksHandler)).Methods("POST")
	}

	if config.C().Mailing.EnableTestMailer {
		r.Handle("/v1/test-email", http.HandlerFunc(testEmailHandler)).Methods("POST")
	}

	c := cors.New(cors.Options{
		AllowedOrigins: config.C().CORS.Origins,
		AllowedHeaders: config.C().CORS.Headers,
		Debug:          true,
	})
	c.Log = log.New(logrus.StandardLogger().Writer(), "[cors] ", log.LstdFlags)
	return handlers.LoggingHandler(logrus.StandardLogger().Writer(), c.Handler(r))
}
