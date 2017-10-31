package dir

import (
	"errors"
	"github.com/SilenMark/util/file"
	"path"
	"path/filepath"
)

type Dir string

// ChildDir only return the dir`s child,And not check the child is exist or is a dir
func (d Dir) ChildDir(child string) Dir {
	Sdir := string(d)
	//path.Clean("/"+child) whill makesure the result can`t access dir`s parrent
	p := filepath.Join(Sdir, filepath.FromSlash(path.Clean("/"+child)))
	return Dir(p)
}
func (d Dir) ToString() string {
	return string(d)
}

// New will check path is exist,and must be a dir
func New(path string) (Dir, error) {
	Sdir, _ := filepath.Abs(path)
	if file.IsDir(Sdir) {
		return Dir(Sdir), nil
	}
	return Dir(""), errors.New(Sdir + "\tNot a Directory Or File Not Exist")

}
