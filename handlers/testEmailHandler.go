package handlers

import (
	"net/http"
)

func testEmailHandler(w http.ResponseWriter, r *http.Request) {
	/*if !cfg.C().Mailing.EnableTestMailer {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var data types.TestEmail
	if !processHandlerData(&data, "testEmail", w, r) {
		return
	}

	if cfg.C().Mailing.EnableTestMailer {
		jobs.MailJobsQueue <- mtypes.MailJob{Kind: mtypes.TestMailJob, Payload: data}
	}*/
}
