package mailing

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/resurtm/boomak-server/cfg"
	"github.com/resurtm/boomak-server/common"
)

type awsCredentialsProvider struct{}

func (m *awsCredentialsProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     cfg.C().Mailing.AccessKeyID,
		SecretAccessKey: cfg.C().Mailing.SecretAccessKey,
		ProviderName:    common.AWSCredentialsProviderName,
	}, nil
}

func (m *awsCredentialsProvider) IsExpired() bool {
	return false
}
