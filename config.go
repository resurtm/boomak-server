package main

import (
	"github.com/jinzhu/configor"
)

const ConfigFileName = "config.yml"

type ConfigType struct {
	Server struct {
		Hostname string `required:"true" default:"localhost"`
		Port     uint   `required:"true" default:"3000"`
	}

	Security struct {
		JwtSigningKey string `required:"true"`
	}

	Database struct {
		User       string `required:"true"`
		Password   string `required:"true"`
		Hostname   string `required:"true" default:"127.0.0.1"`
		Port       uint   `required:"true" default:"27017"`
		Name       string `required:"true"`
		NoPassword bool   `required:"true" default:"false"`
	}

	Cors struct {
		Origins []string ``
		Debug   bool     `required:"true" default:"false"`
	}

	Mailing struct {
		AccessKeyID     string `required:"true"`
		SecretAccessKey string `required:"true"`
		AwsRegion       string `required:"true"`

		EnableTests bool `required:"true" default:"false"`
		Debug       bool `required:"true" default:"false"`

		SignupWorkers  uint `required:"true" default:"2"`
		SignupChanSize uint `required:"true" default:"255"`
	}
}

var Config ConfigType

func LoadConfig() ConfigType {
	if err := configor.Load(&Config, ConfigFileName); err != nil {
		panic(err)
	}
	return Config
}
