package config

import (
	"fmt"
)

type configuration struct {
	Server struct {
		Hostname string `required:"true"`
		Port     uint   `required:"true"`
	}

	Security struct {
		JWTSigningKey string `required:"true"`

		EnableFaker bool ``
	}

	Database struct {
		User       string ``
		Password   string ``
		Hostname   string `default:"localhost"`
		Port       uint   `default:"27017"`
		Name       string `required:"true"`
		NoPassword bool   ``
	}

	CORS struct {
		Origins []string ``
		Headers []string ``
	}

	Mailing struct {
		AccessKeyID     string `required:"true"`
		SecretAccessKey string `required:"true"`
		AWSRegion       string `required:"true"`

		FromEmail               string `required:"true"`
		VerificationTokenLength byte   `default:"8"`
		EnableTestMailer        bool   ``

		WorkerCount     byte `default:"2"`
		WorkerQueueSize uint `default:"255"`
	}
}

func (config configuration) ListenAddr() string {
	return fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port)
}

func (config configuration) DSN() (dsn string) {
	db := config.Database
	if db.NoPassword {
		dsn = fmt.Sprintf("mongodb://%s:%d/%s", db.Hostname, db.Port, db.Name)
	} else {
		// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", db.User, db.Password, db.Hostname, db.Port, db.Name)
	}
	return
}
