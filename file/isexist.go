package file

import (
	"os"
)

// IsExist if return true the filename is exist,else not exist
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
