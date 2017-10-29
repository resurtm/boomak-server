package config

import "github.com/jinzhu/configor"

type configType struct {
	Server struct {
		Hostname string `required:"true"`
		Port     uint   `required:"true"`
	}

	Security struct {
		JWTSigningKey string `required:"true"`
	}

	Database struct {
		User       string ``
		Password   string ``
		Hostname   string `default:"localhost"`
		Port       uint   `default:"27017"`
		Name       string `required:"true"`
		NoPassword bool   ``
	}

	Cors struct {
		Origins []string ``
		Debug   bool     ``
	}

	Mailing struct {
		AccessKeyID     string `required:"true"`
		SecretAccessKey string `required:"true"`
		AWSRegion       string `required:"true"`

		EnableTestMailer bool ``
		Debug            bool ``

		WorkerCount     uint `default:"2"`
		WorkerQueueSize uint `default:"255"`
	}
}

var configData configType

const configFileName = "config.yml"

func init() {
	if err := configor.Load(&configData, configFileName); err != nil {
		panic(err)
	}
}

func Config() configType {
	return configData
}
