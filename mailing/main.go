package mailing

import (
	"github.com/resurtm/boomak-server/config"
	. "github.com/resurtm/boomak-server/mailing/base"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	tjses "github.com/tj/go-ses"
	"github.com/aws/aws-sdk-go/service/ses"
)

func init() {
	MailJobsQueue = make(chan MailJob, config.C().Mailing.WorkerQueueSize)
	for i := byte(1); i <= config.C().Mailing.WorkerCount; i++ {
		go mailJobsWorker(i, MailJobsQueue)
	}
}

func newClient() *tjses.Client {
	creds := credentials.NewCredentials(&awsCredentialsProvider{})
	cfg := aws.NewConfig().WithRegion(config.C().Mailing.AWSRegion).WithCredentials(creds)
	return tjses.New(ses.New(session.New(cfg)))
}
