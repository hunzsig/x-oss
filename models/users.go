package models

import (
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Token          string `gorm:"primary_key;type:char(255)"` // 用户唯一token key
	Status         uint8                                      // 状态 -1无效 1有效
	Level          uint32                                     // 等级
	Exp            uint32                                     // 经验
	AllowSizeUnit  string `gorm:"type:char(255)"`             // 文件大小单位
	AllowSizeTotal int64                                      // 允许储存的总体文件大小
	AllowSizeOne   int64                                      // 允许储存的单个文件大小
	AllowQty       int64                                      // 允许储存的文件数
	AllowSuffix    string `gorm:"type:varchar(1024)"`         // 允许的文件后缀
}
