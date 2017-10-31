package handlers

import (
	"net/http"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/mailing"
	"github.com/resurtm/boomak-server/types"
)

func testEmailHandler(w http.ResponseWriter, r *http.Request) {
	if !cfg.C().Mailing.EnableTestMailer {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var data types.TestEmail
	if !processHandlerData(&data, "testEmail", w, r) {
		return
	}

	mailing.SendTestEmail(data.RecipientEmail, data.TestString)
}
