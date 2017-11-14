package handlers_old

import (
	"net/http"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/mailing/jobs"
	mtypes "github.com/resurtm/boomak-server/mailing/types"
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

	if cfg.C().Mailing.EnableTestMailer {
		jobs.MailJobsQueue <- mtypes.MailJob{Kind: mtypes.TestMailJob, Payload: data}
	}
}
