package config

import (
	"github.com/jinzhu/configor"
	"flag"
	log "github.com/sirupsen/logrus"
)

const mainConfigDefaultFileName = "config.yml"

var mainConfig configuration

func C() configuration {
	return mainConfig
}

func init() {
	var mainConfigFileName string
	flag.StringVar(&mainConfigFileName, "config", mainConfigDefaultFileName, "configuration file name")
	flag.Parse()

	log.Infof("using configuration from file \"%s\"", mainConfigFileName)
	if err := configor.Load(&mainConfig, mainConfigFileName); err != nil {
		panic(err)
	}
}
