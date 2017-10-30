package mailing

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
	tj "github.com/tj/go-ses"
	"github.com/resurtm/boomak-server/database"
	cfg "github.com/resurtm/boomak-server/config"
)

const (
	mailJobTest       = iota
	mailJobSignup     = iota
	mailJobRegistered = iota
)

type mailJob struct {
	kind    byte
	payload interface{}
}

type mailJobTestType struct {
	recipient string
	data string
}

var mailJobsQueue chan mailJob

func init() {
	mailJobsQueue = make(chan mailJob, cfg.Config().Mailing.WorkerQueueSize)
	for i := byte(1); i <= cfg.Config().Mailing.WorkerCount; i++ {
		go mailWorker(i, mailJobsQueue)
	}
}

func newClient() *tj.Client {
	creds := credentials.NewCredentials(&awsCredsProvider{})
	config := aws.NewConfig().
		WithRegion(cfg.Config().Mailing.AWSRegion).
		WithCredentials(creds)
	return tj.New(ses.New(session.New(config)))
}

func SendTestEmail(recipientEmail string, str string) {
	if cfg.Config().Mailing.EnableTestMailing {
		mailJobsQueue <- mailJob{
			kind:    mailJobTest,
			payload: mailJobTestType{recipient: recipientEmail, data: str},
		}
	}
}

func SendSignupEmail(user database.User) {
	mailJobsQueue <- mailJob{kind: mailJobSignup, payload: user}
}

func SendRegisteredEmail(user database.User) {
	mailJobsQueue <- mailJob{kind: mailJobRegistered, payload: user}
}
