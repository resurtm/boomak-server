package main

import (
	"net/http"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/handlers"
)

func main() {
	http.ListenAndServe(
		cfg.C().ListenAddr(),
		handlers.New(),
	)
}
