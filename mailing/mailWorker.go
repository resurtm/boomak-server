package mailing

import (
	"fmt"
	tj "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/database"
	cfg "github.com/resurtm/boomak-server/config"
	"bytes"
	"html/template"
	"github.com/resurtm/boomak-server/tools"
	"path/filepath"
)

func mailWorker(workerID byte, ch <-chan mailJob) {
	if cfg.Config().Mailing.Debug {
		fmt.Printf("Mail worker #%d: started!\n", workerID)
	}

	client := newClient()
	var email tj.Email

	for {
		job := <-ch

		if !cfg.Config().Mailing.EnableTestMailing && job.kind == mailJobTest {
			continue
		}

		if cfg.Config().Mailing.Debug {
			fmt.Printf("Mail worker #%d: started processing mail job type #%d\n", workerID, job.kind)
		}

		switch job.kind {
		case mailJobTest:
			email = prepareTestEmail(job.payload.(mailJobTestType))
		case mailJobSignup:
			email = prepareSignupEmail(job.payload.(database.User))
		case mailJobRegistered:
			email = prepareRegisteredEmail(job.payload.(database.User))
		}

		if err := client.SendEmail(email); err != nil {
			fmt.Printf("Mail worker #%d: failure: %s\n", workerID, err.Error())
		}
		if cfg.Config().Mailing.Debug {
			fmt.Printf("Mail worker #%d: finished processing mail job type #%d\n", workerID, job.kind)
		}
	}
}

func prepareTestEmail(job mailJobTestType) tj.Email {
	return tj.Email{
		From:    cfg.Config().Mailing.FromEmail,
		To:      []string{job.recipient},
		Subject: "Boomak Test Email",
		Text:    fmt.Sprintf("Boomak Test Email\n\nTest String: %s", job.data),
		HTML:    fmt.Sprintf("<h1>Boomak Test Email</h1><p>Test String: %s</p>", job.data),
	}
}

func prepareSignupEmail(user database.User) tj.Email {
	t, err := template.ParseFiles(filepath.Join(tools.CurrentDir(), "templates", "mailJobSignup.html"))
	if err != nil {
		panic(err)
	}

	data := struct {
		Key string
	}{
		Key: "test string",
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		panic(err)
	}

	return tj.Email{
		From:    cfg.Config().Mailing.FromEmail,
		To:      []string{user.Email},
		Subject: "Welcome to Boomak!",
		Text:    "Welcome to Boomak!",
		HTML:    tpl.String(),
	}
}

func prepareRegisteredEmail(user database.User) tj.Email {
	return tj.Email{}
}
