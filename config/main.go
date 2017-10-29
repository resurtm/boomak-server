package config

import "github.com/jinzhu/configor"

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
