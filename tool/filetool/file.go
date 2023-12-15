package filetool

import "os"

func Create(name string) (f *os.File, err error) {
	err = os.MkdirAll(name, os.ModeDir)
	if err != nil {
		return
	}
	return os.Create(name)
}
