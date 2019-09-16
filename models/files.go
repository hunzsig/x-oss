package models

import (
	"time"
)

type Files struct {
	Hash         string `gorm:"primary_key;type:char(255);not null;unique"` // hash
	Name         string `gorm:"type:char(255)"`                             // 文件名
	TokenName    string `gorm:"type:char(255)"`                             // token名
	Suffix       string `gorm:"type:char(255)"`                             // 文件后缀
	Size         uint64                                                     // 文件大小
	ContentType  string `gorm:"type:char(255)"`                             // 内容类型
	Path         string `gorm:"type:char(255)"`                             // 存在路径
	Uri          string `gorm:"type:char(255)"`                             // URI
	FromUrl      string `gorm:"type:char(255)"`                             // 来源地址
	CallQty      uint32                                                     // 调用次数
	CallLastTime *time.Time                                                 // 最后一次调用时间
	CreateTime   *time.Time                                                 // 创建日期
	UpdateTime   *time.Time                                                 // 更新日期
}
