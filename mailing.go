package main

var SignupMails chan User

func InitMailing() {
	InitAwsSes()

	SignupMails = make(chan User, Config.Mailing.SignupChanSize)
	for i := uint(1); i <= Config.Mailing.SignupWorkers; i++ {
		go SignupMailSender(i, SignupMails)
	}
}
