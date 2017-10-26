package main

import "github.com/aws/aws-sdk-go/aws/credentials"

type AwsCredsProvider struct{}

func (m *AwsCredsProvider) Retrieve() (credentials.Value, error) {
	return credentials.Value{
		AccessKeyID:     Config.Mailing.AccessKeyID,
		SecretAccessKey: Config.Mailing.SecretAccessKey,
		ProviderName:    "AwsCredsProvider",
	}, nil
}

func (m *AwsCredsProvider) IsExpired() bool {
	return false
}
