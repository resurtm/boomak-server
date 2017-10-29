package config

import "fmt"

type configType struct {
	Server struct {
		Hostname    string `required:"true"`
		Port        uint   `required:"true"`
		DebugOutput bool   ``
	}

	Security struct {
		JWTSigningKey string `required:"true"`
		JSONSchemaDir string `default:"jsonSchema"`
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

func (ct configType) ListenAddr() string {
	return fmt.Sprintf("%s:%d", ct.Server.Hostname, ct.Server.Port)
}