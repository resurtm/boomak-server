package cfg

import "fmt"

type configuration struct {
	Server struct {
		Hostname string `required:"true"`
		Port     uint   `required:"true"`
		Debug    bool   ``
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

	CORS struct {
		Origins []string ``
		Headers []string ``
		Debug   bool     ``
	}

	Mailing struct {
		AccessKeyID     string `required:"true"`
		SecretAccessKey string `required:"true"`
		AWSRegion       string `required:"true"`

		FromEmail               string `required:"true"`
		VerificationTokenLength byte   `default:"8"`
		EnableTestMailer        bool   ``
		Debug                   bool   ``

		WorkerCount      byte   `default:"2"`
		WorkerQueueSize  uint   `default:"255"`
	}
}

func (c configuration) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Hostname, c.Server.Port)
}

func (c configuration) DSN() (dsn string) {
	d := c.Database
	if d.NoPassword {
		dsn = fmt.Sprintf("mongodb://%s:%d/%s", d.Hostname, d.Port, d.Name)
	} else {
		// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
		dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", d.User, d.Password, d.Hostname, d.Port, d.Name)
	}
	return
}
