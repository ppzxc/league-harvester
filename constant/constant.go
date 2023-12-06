package constant

import (
	"os"
	"path/filepath"
)

var RoamingHome string
var LoggingHome string

func LoggingFilePath(fileName string) string {
	return filepath.Join(LoggingHome, fileName)
}

func init() {
	dir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	RoamingHome = filepath.Join(dir, "Baecorp", "Havester")
	LoggingHome = filepath.Join(RoamingHome, "log")
}
