package system

import (
	"time"
)

var systemStart int64

func Now() int64 {
	return time.Now().UnixNano()
}

func Start() {
	systemStart = time.Now().UnixNano()
}

func End() int64 {
	return (time.Now().UnixNano() - systemStart) / 1000
}
