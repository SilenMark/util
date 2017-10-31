package dir

import (
	"errors"
	"github.com/SilenMark/util/file"
	"path"
	"path/filepath"
)

type dir string

// ChildDir only return the dir`s child,And not check the child is exist or is a dir
func (d dir) ChildDir(child string) dir {
	Sdir := string(d)
	//path.Clean("/"+child) whill makesure the result can`t access dir`s parrent
	p := filepath.Join(Sdir, filepath.FromSlash(path.Clean("/"+child)))
	return dir(p)
}
func (d dir) ToString() string {
	return string(d)
}

// New will check path is exist,and must be a dir
func New(path string) (dir, error) {
	Sdir, _ := filepath.Abs(path)
	if file.IsDir(Sdir) {
		return dir(Sdir), nil
	}
	return dir(""), errors.New(Sdir + "\tNot a Directory Or File Not Exist")

}
