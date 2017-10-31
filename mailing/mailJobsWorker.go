package mailing

import (
	"fmt"
	tjses "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/user"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/common"
	mtypes "github.com/resurtm/boomak-server/mailing/types"
	"github.com/resurtm/boomak-server/types"
)

func mailJobsWorker(workerID byte, ch <-chan mtypes.MailJob) {
	if cfg.C().Mailing.Debug {
		fmt.Printf("Mail worker #%d: started!\n", workerID)
	}

	client := newClient()
	for {
		job := <-ch
		if !cfg.C().Mailing.EnableTestMailer && job.Kind == mtypes.TestMailJob {
			continue
		}
		if cfg.C().Mailing.Debug {
			fmt.Printf("Mail worker #%d: started processing mail job type #%d\n", workerID, job.Kind)
		}

		email := emailBuilders[job.Kind](job.Payload)
		if cfg.C().Mailing.Debug {
			fmt.Printf("Mail worker #%d: %#v\n", workerID, email)
		}

		if err := client.SendEmail(email); err != nil {
			fmt.Printf("Mail worker #%d: failure: %s\n", workerID, err.Error())
		}
		if cfg.C().Mailing.Debug {
			fmt.Printf("Mail worker #%d: finished processing mail job type #%d\n", workerID, job.Kind)
		}
	}
}

var emailBuilders = map[byte]func(interface{}) tjses.Email{
	// email builder for type testMailJob
	mtypes.TestMailJob: func(raw interface{}) tjses.Email {
		if job, ok := raw.(types.TestEmail); !ok {
			panic("Invalid job payload has been provided")
		} else {
			return tjses.Email{
				From:    cfg.C().Mailing.FromEmail,
				To:      []string{job.RecipientEmail},
				Subject: common.TestEmailSubject,
				Text:    fmt.Sprintf("Boomak Test Email\n\nTest String: %s", job.TestString),
				HTML:    fmt.Sprintf("<h1>Boomak Test Email</h1><p>Test String: %s</p>", job.TestString),
			}
		}
	},

	// email builder for type emailVerifyMailJob
	mtypes.EmailVerifyMailJob: func(raw interface{}) tjses.Email {
		u, ok := raw.(user.User)
		if !ok {
			panic("Invalid job payload has been provided")
		}

		tplData := struct {
			User user.User
		}{User: u}

		return tjses.Email{
			From:    cfg.C().Mailing.FromEmail,
			To:      []string{u.Email},
			Subject: common.VerifyEmailSubject,
			Text:    renderTextTemplate("emailVerifyMail", tplData),
			HTML:    renderHtmlTemplate("emailVerifyMail", tplData),
		}
	},

	// email builder for type signupFinishedMailJob
	mtypes.SignupFinishedMailJob: func(raw interface{}) tjses.Email {
		// todo: write code here
		return tjses.Email{}
	},
}
