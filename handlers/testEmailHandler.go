package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/config"
	"github.com/resurtm/boomak-server/common"
	mailing "github.com/resurtm/boomak-server/mailing/base"
	log "github.com/sirupsen/logrus"
)

func testEmailHandler(w http.ResponseWriter, r *http.Request) {
	if !config.C().Mailing.EnableTestMailer {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("you're trying to use forbidden endpoint"))
		log.Warn("tried to use forbidden/disabled endpoint")
		return
	}

	var data common.TestEmail
	if !processHandlerData(&data, "testEmail", w, r) {
		return
	}

	mailing.EnqueueTestMailJob(data)
}
