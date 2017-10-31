package main

import (
	"net/http"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/handlers"
	"github.com/resurtm/boomak-server/mailing"
)

func main() {
	mailing.InitMailing()
	http.ListenAndServe(
		cfg.C().ListenAddr(),
		handlers.New(),
	)
}
