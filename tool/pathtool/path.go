package pathtool

import (
	"path/filepath"
	"runtime"
)

var (
	projectRoot string
)

func init() {
	_, b, _, _ := runtime.Caller(0)
	projectRoot = filepath.Join(filepath.Dir(b), "..")
}

func ProjectRoot() string {
	return projectRoot
}
