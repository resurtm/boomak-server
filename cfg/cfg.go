package cfg

import (
	"github.com/jinzhu/configor"
	"github.com/resurtm/boomak-server/common"
)

var mainConfig configuration

func C() configuration {
	return mainConfig
}

func init() {
	if err := configor.Load(&mainConfig, common.MainConfigFileName); err != nil {
		panic(err)
	}
}
