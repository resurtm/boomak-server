package cfg

import (
	"github.com/jinzhu/configor"
)

const mainConfigFileName = "config.yml"

var mainConfig configuration

func C() configuration {
	return mainConfig
}

func init() {
	if err := configor.Load(&mainConfig, mainConfigFileName); err != nil {
		panic(err)
	}
}
