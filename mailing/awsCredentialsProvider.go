package mailing

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/resurtm/boomak-server/config"
)

type awsCredentialsProvider struct{}

func (m *awsCredentialsProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     config.C().Mailing.AccessKeyID,
		SecretAccessKey: config.C().Mailing.SecretAccessKey,
		ProviderName:    "boomak",
	}, nil
}

func (m *awsCredentialsProvider) IsExpired() bool {
	return false
}
