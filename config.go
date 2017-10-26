package main

import (
	"github.com/jinzhu/configor"
	"fmt"
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
		Hostname   string `required:"true" default:"localhost"`
		Port       uint   `required:"true" default:"27017"`
		Name       string `required:"true"`
		NoPassword bool   `required:"true" default:"false"`
	}
	Cors struct {
		Origins []string ``
		Debug   bool     `required:"true" default:"false"`
	}
}

var Config ConfigType

func LoadConfig() ConfigType {
	configor.Load(&Config, ConfigFileName)
	fmt.Println(Config.Cors)
	return Config
}
