package mailing

import (
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/user"
)

// enum consts for mailJob struct
const (
	testMailJob           = iota
	emailVerifyMailJob    = iota
	signupFinishedMailJob = iota
)

type mailJob struct {
	kind    byte
	payload interface{}
}

type testMailJobPayload struct {
	recipient string
	data      string
}

var mailJobsQueue chan mailJob

func init() {
	mailJobsQueue = make(chan mailJob, cfg.C().Mailing.WorkerQueueSize)
	for i := byte(1); i <= cfg.C().Mailing.WorkerCount; i++ {
		go mailJobsWorker(i, mailJobsQueue)
	}
}

func SendTestEmail(recipientEmail string, str string) {
	if cfg.C().Mailing.EnableTestMailer {
		mailJobsQueue <- mailJob{
			kind:    testMailJob,
			payload: testMailJobPayload{recipient: recipientEmail, data: str},
		}
	}
}

func SendEmailVerifyEmail(user *user.User) {
	mailJobsQueue <- mailJob{kind: emailVerifyMailJob, payload: user}
}

func SendSignupFinishedEmail(user *user.User) {
	mailJobsQueue <- mailJob{kind: signupFinishedMailJob, payload: user}
}
