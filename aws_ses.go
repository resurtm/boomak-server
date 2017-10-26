package main

import (
	tj "github.com/tj/go-ses"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var AwsSesClient *tj.Client

func InitAwsSes() {
	creds := credentials.NewCredentials(&AwsCredsProvider{})

	config := aws.NewConfig().
		WithRegion("eu-west-1").
		WithCredentials(creds)

	AwsSesClient = tj.New(ses.New(session.New(config)))
}
