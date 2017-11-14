package main

import (
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/handlers"
	"net/http"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("app started")

	//mailing.InitMailing()
	http.ListenAndServe(config.C().ListenAddr(), handlers.New())

	log.Info("app exited")
}
