package filetool

import (
	"os"
	"path/filepath"
)

func Open(name string, append bool) (f *os.File, err error) {
	err = os.MkdirAll(filepath.Dir(name), os.ModeDir|0755)
	if err != nil {
		return
	}
	flag := os.O_RDWR | os.O_CREATE
	if append {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}
	return os.OpenFile(name, flag, 0666)
}
