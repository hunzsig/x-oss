package system

import (
	godump "github.com/favframework/debug"
	"os"
	"path/filepath"
)

func Dump(data interface{}) {
	godump.Dump(data)
}

func CurrentPath() string {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return path
}

