package mailing

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/aws/aws-sdk-go/aws/session"
	tjses "github.com/tj/go-ses"
	"github.com/aws/aws-sdk-go/service/ses"
)

func newClient() *tjses.Client {
	creds := credentials.NewCredentials(&awsCredentialsProvider{})
	config := aws.NewConfig().
		WithRegion(cfg.C().Mailing.AWSRegion).
		WithCredentials(creds)
	return tjses.New(ses.New(session.New(config)))
}
