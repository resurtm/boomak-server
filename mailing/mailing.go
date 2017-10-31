package mailing

import (
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/mailing/types"
	"github.com/resurtm/boomak-server/mailing/jobs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	tjses "github.com/tj/go-ses"
	"github.com/aws/aws-sdk-go/service/ses"
)

func InitMailing() {
	jobs.MailJobsQueue = make(chan types.MailJob, cfg.C().Mailing.WorkerQueueSize)
	for i := byte(1); i <= cfg.C().Mailing.WorkerCount; i++ {
		go mailJobsWorker(i, jobs.MailJobsQueue)
	}
}

func newClient() *tjses.Client {
	creds := credentials.NewCredentials(&awsCredentialsProvider{})
	config := aws.NewConfig().
		WithRegion(cfg.C().Mailing.AWSRegion).
		WithCredentials(creds)
	return tjses.New(ses.New(session.New(config)))
}
