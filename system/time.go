package system

import (
	"../php2go"
)

var systemStart int64

func Start() {
	systemStart = php2go.Microtime()
}

func End() int64 {
	return (php2go.Microtime() - systemStart) / 1000
}
