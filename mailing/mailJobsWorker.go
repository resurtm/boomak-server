package mailing

import (
	"fmt"
	tjses "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/user"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/common"
)

func mailJobsWorker(workerID byte, ch <-chan mailJob) {
	if cfg.C().Mailing.Debug {
		fmt.Printf("Mail worker #%d: started!\n", workerID)
	}

	client := newClient()
	var email tjses.Email

	for {
		job := <-ch

		if !cfg.C().Mailing.EnableTestMailer && job.kind == testMailJob {
			continue
		}
		if cfg.C().Mailing.Debug {
			fmt.Printf("Mail worker #%d: started processing mail job type #%d\n", workerID, job.kind)
		}

		emailBuilders[job.kind](job.payload)

		if err := client.SendEmail(email); err != nil {
			fmt.Printf("Mail worker #%d: failure: %s\n", workerID, err.Error())
		}
		if cfg.C().Mailing.Debug {
			fmt.Printf("Mail worker #%d: finished processing mail job type #%d\n", workerID, job.kind)
		}
	}
}

var emailBuilders = map[byte]func(interface{}) tjses.Email{
	// email builder for type testMailJob
	testMailJob: func(raw interface{}) tjses.Email {
		if job, ok := raw.(testMailJobPayload); !ok {
			panic("Invalid job payload has been provided")
		} else {
			return tjses.Email{
				From:    cfg.C().Mailing.FromEmail,
				To:      []string{job.recipient},
				Subject: common.TestEmailSubject,
				Text:    fmt.Sprintf("Boomak Test Email\n\nTest String: %s", job.data),
				HTML:    fmt.Sprintf("<h1>Boomak Test Email</h1><p>Test String: %s</p>", job.data),
			}
		}
	},

	// email builder for type emailVerifyMailJob
	emailVerifyMailJob: func(raw interface{}) tjses.Email {
		user, ok := raw.(user.User)
		if !ok {
			panic("Invalid job payload has been provided")
		}

		tplData := struct {
			User user.User
		}{User: user}

		return tjses.Email{
			From:    cfg.C().Mailing.FromEmail,
			To:      []string{user.Email},
			Subject: common.VerifyEmailSubject,
			Text:    renderTextTemplate("emailVerifyMail", tplData),
			HTML:    renderHtmlTemplate("emailVerifyMail", tplData),
		}
	},

	// email builder for type signupFinishedMailJob
	signupFinishedMailJob: func(raw interface{}) tjses.Email {
		// todo: write code here
		return tjses.Email{}
	},
}
