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

func GetFileContent(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		file.Close()
	}()
	b := make([]byte, 4096)
	n, err := file.Read(b)
	if err != nil {
		return "", err
	}
	data := string(b[:n])
	return data, nil
}
