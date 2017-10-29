package main

import (
	"net/http"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/handlers"
)

func main() {
	http.ListenAndServe(
		config.Config().ListenAddr(),
		handlers.New(),
	)
}
