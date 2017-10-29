package mailing

import (
	"fmt"
	tj "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/database"
	cfg "github.com/resurtm/boomak-server/config"
)

func mailWorker(workerID byte, ch <-chan database.User) {
	if cfg.Config().Mailing.Debug {
		fmt.Printf("Mail worker #%d: started!\n", workerID)
	}

	client := newClient()

	for {
		user := <-ch
		if cfg.Config().Mailing.Debug {
			fmt.Printf("Mail worker #%d: sending email to user #%d\n", workerID, user.Id)
		}

		email := tj.Email{
			From:    "resurtm@gmail.com",
			To:      []string{user.Email},
			Subject: "Welcome to Boomak!",
			Text:    "Welcome to Boomak!",
			HTML:    "<h1>Welcome to Boomak!</h1>",
		}

		if err := client.SendEmail(email); err != nil {
			fmt.Printf("Mail worker #%d: failure: %s\n", workerID, err.Error())
		}

		if cfg.Config().Mailing.Debug {
			fmt.Printf("Mail worker #%d: email has been to user #%d\n", workerID, user.Id)
		}
	}
}
