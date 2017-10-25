package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	loadConfig()
	connectToDb()
	initHttp()
}

func initHttp() {
	router := mux.NewRouter()

	router.Handle("/auth", http.HandlerFunc(authHandler)).Methods("GET")
	router.Handle("/register", http.HandlerFunc(registerHandler)).Methods("POST")

	listenAddr := fmt.Sprintf(":%d", config.Server.Port)
	fmt.Printf("Listening at \"%s\"...", listenAddr)
	http.ListenAndServe(listenAddr, handlers.LoggingHandler(os.Stdout, router))
}
