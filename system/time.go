package system

import (
	"time"
)

var systemStart int64

func MicroTime() int64 {
	return time.Now().UnixNano()
}

func Start() {
	systemStart = MicroTime()
}

func End() int64 {
	return (MicroTime() - systemStart) / 1000
}
