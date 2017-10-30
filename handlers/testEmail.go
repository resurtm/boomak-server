package handlers

import (
	"net/http"
	cfg "github.com/resurtm/boomak-server/config"
	"encoding/json"
	"github.com/resurtm/boomak-server/mailing"
)

func testEmailHandler(w http.ResponseWriter, r *http.Request) {
	if !cfg.Config().Mailing.EnableTestMailing {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validateHandlerData(data, "testEmail", w) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	mailing.SendTestEmail(data["recipientEmail"].(string), data["testString"].(string))
}
