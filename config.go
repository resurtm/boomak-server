package main

import "github.com/jinzhu/configor"

type configType struct {
	Server struct {
		Port uint `required:"true" default:"8000"`
	}
	Security struct {
		JwtSigningKey string `required:"true"`
	}
}

const configFileName = "config.yml"

var config configType

func loadConfig() configType {
	configor.Load(&config, configFileName)
	return config
}
