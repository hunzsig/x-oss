package models

import (
	"github.com/jinzhu/gorm"
)

type Setting struct {
	gorm.Model
	Key  string `gorm:"primary_key;type:char(255);not null;unique"` // key
	Data string `gorm:"type:varchar(1024)"`                         // 文件名
}
