package system

import (
	"github.com/favframework/debug"
)

func Dump(data interface{}) {
	godump.Dump(data)
}
