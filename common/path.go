package common

import (
	"path/filepath"
	"os"
)

func CurrentDir() string {
	if currDir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	} else {
		return currDir
	}
}
