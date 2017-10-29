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

var signupMailPayload chan database.User

func newClient() *tj.Client {
	creds := credentials.NewCredentials(&awsCredsProvider{})
	config := aws.NewConfig().
		WithRegion(cfg.Config().Mailing.AWSRegion).
		WithCredentials(creds)
	return tj.New(ses.New(session.New(config)))
}

func SendSignupEmail(user database.User) {
	signupMailPayload <- user
}

func init() {
	signupMailPayload = make(chan database.User, cfg.Config().Mailing.WorkerQueueSize)

	for i := byte(1); i <= cfg.Config().Mailing.WorkerCount; i++ {
		go mailWorker(i, signupMailPayload)
	}
}
