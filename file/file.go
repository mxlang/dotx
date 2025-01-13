package file

import (
	"os"
)

func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		if fileInfo.IsDir() {
			return true
		}
	}

	return false
}
