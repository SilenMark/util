package file

import (
	"os"
)

func IsDir(path string) bool {
	p, _ := os.Stat(path)
	if p != nil {
		return p.IsDir()
	}
	return false
}
