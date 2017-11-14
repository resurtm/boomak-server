package mailing

import (
	"fmt"
	tjses "github.com/tj/go-ses"
	log "github.com/sirupsen/logrus"
	"github.com/resurtm/boomak-server/user"
	"github.com/resurtm/boomak-server/common"
	"github.com/resurtm/boomak-server/config"
	. "github.com/resurtm/boomak-server/mailing/base"
)

var emailBuilders = map[byte]func(interface{}) *tjses.Email{
	TestMailJob: func(raw interface{}) *tjses.Email {
		if job, ok := raw.(common.TestEmail); !ok {
			log.WithField("raw", raw).Warn("invalid job payload has been provided")
			return nil
		} else {
			return &tjses.Email{
				From:    config.C().Mailing.FromEmail,
				To:      []string{job.RecipientEmail},
				Subject: "Boomak: Test Email",
				Text:    fmt.Sprintf("Boomak Test Email\n\nTest String: %s", job.TestString),
				HTML:    fmt.Sprintf("<h1>Boomak Test Email</h1><p>Test String: %s</p>", job.TestString),
			}
		}
	},

	EmailVerifyMailJob: func(raw interface{}) *tjses.Email {
		usr, ok := raw.(user.User)
		if !ok {
			log.WithField("raw", raw).Warn("invalid job payload has been provided")
			return nil
		}

		tplData := struct {
			User user.User
		}{User: usr}
		return &tjses.Email{
			From:    config.C().Mailing.FromEmail,
			To:      []string{usr.Email},
			Subject: "Boomak: Confirm Your Email",
			Text:    renderTextTemplate("emailVerifyMail", tplData),
			HTML:    renderHtmlTemplate("emailVerifyMail", tplData),
		}
	},

	SignupFinishedMailJob: func(raw interface{}) *tjses.Email {
		usr, ok := raw.(user.User)
		if !ok {
			log.WithField("raw", raw).Warn("invalid job payload has been provided")
			return nil
		}

		tplData := struct {
			User user.User
		}{User: usr}
		return &tjses.Email{
			From:    config.C().Mailing.FromEmail,
			To:      []string{usr.Email},
			Subject: "Boomak: Welcome to Application!",
			Text:    renderTextTemplate("signupFinishedMail", tplData),
			HTML:    renderHtmlTemplate("signupFinishedMail", tplData),
		}
	},
}
