package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/handlers"
	_ "github.com/resurtm/boomak-server/mailing"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Info("app started")
	http.ListenAndServe(config.C().ListenAddr(), handlers.New())
	log.Info("app exited")
}
