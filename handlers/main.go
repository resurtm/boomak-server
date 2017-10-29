package handlers

import (
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
	"github.com/resurtm/boomak-server/config"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.Handle("/auth", http.HandlerFunc(authHandler)).Methods("POST")
	r.Handle("/register", http.HandlerFunc(signupHandler)).Methods("POST")
	r.Handle("/validate", http.HandlerFunc(validateHandler)).Methods("POST")

	h := setupCORS(r)
	if config.Config().Server.DebugOutput {
		h = handlers.LoggingHandler(os.Stdout, h)
	}
	return h
}

func setupCORS(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: config.Config().CORS.Origins,
		Debug:          config.Config().CORS.Debug,
	})
	return c.Handler(handler)
}
