package system

import (
	"github.com/favframework/debug"
	"time"
)

var systemStart int64

func Start() {
	systemStart = time.Now().UnixNano()
}

func End() int64 {
	return (time.Now().UnixNano() - systemStart) / 1000
}

func Dump(data interface{}) {
	godump.Dump(data)
}
