package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func main() {
	LoadConfig()
	ConnectToDb()
	InitHttp()
}

func InitHttp() {
	router := mux.NewRouter()

	router.Handle("/auth", http.HandlerFunc(AuthHandler)).Methods("POST")
	router.Handle("/register", http.HandlerFunc(RegisterHandler)).Methods("POST")

	listenAddr := fmt.Sprintf("%s:%d", Config.Server.Hostname, Config.Server.Port)
	fmt.Printf("Listening at \"%s\"...\n", listenAddr)
	http.ListenAndServe(listenAddr, handlers.LoggingHandler(os.Stdout, router))
}
