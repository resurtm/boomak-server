package main

import (
	"fmt"
	"time"
	tj "github.com/tj/go-ses"
)

func SignupMailSender(wrkId uint, ch <-chan User) {
	if Config.Mailing.Debug {
		fmt.Printf("SignupMailSender #%d: started!\n", wrkId)
	}
	for {
		user := <-ch
		if Config.Mailing.Debug {
			fmt.Printf("SignupMailSender #%d: sending email to user #%d\n", wrkId, user.Id)
		}
		time.Sleep(time.Second * 1)

		email := tj.Email{
			From:    "resurtm@gmail.com",
			To:      []string{user.Email},
			Subject: "Welcome to Boomak!",
			Text:    "Welcome to Boomak!",
			HTML:    "<h1>Welcome to Boomak!</h1>",
		}
		if err := AwsSesClient.SendEmail(email); err != nil {
			fmt.Printf("SignupMailSender #%d: failure: %s\n", wrkId, err.Error())
		}

		time.Sleep(time.Second * 1)
		if Config.Mailing.Debug {
			fmt.Printf("SignupMailSender #%d: email sent to user #%d\n", wrkId, user.Id)
		}
	}
}
