package mailing

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/resurtm/boomak-server/config"
)

type awsCredsProvider struct{}

func (m *awsCredsProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     config.Config().Mailing.AccessKeyID,
		SecretAccessKey: config.Config().Mailing.SecretAccessKey,
		ProviderName:    "awsCredsProvider",
	}, nil
}

func (m *awsCredsProvider) IsExpired() bool {
	return false
}
