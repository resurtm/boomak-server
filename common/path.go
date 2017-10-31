package common

import (
	"path/filepath"
	"os"
)

func CurrentDir() string {
	currDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return currDir
}
