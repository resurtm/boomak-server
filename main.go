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
	InitMailing()
	InitHttp()
}

func InitHttp() {
	r := mux.NewRouter()

	r.Handle("/auth", http.HandlerFunc(AuthHandler)).Methods("POST")
	r.Handle("/register", http.HandlerFunc(RegisterHandler)).Methods("POST")
	r.Handle("/validate", http.HandlerFunc(ValidateHandler)).Methods("POST")

	listenAddr := fmt.Sprintf("%s:%d", Config.Server.Hostname, Config.Server.Port)
	fmt.Printf("Listening at \"%s\"...\n", listenAddr)

	h1 := SetupCors(r)
	h2 := handlers.LoggingHandler(os.Stdout, h1)
	http.ListenAndServe(listenAddr, h2)
}
