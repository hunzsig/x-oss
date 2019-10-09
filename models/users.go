package models

type Users struct {
	Token         string `gorm:"type:char(255);primary_key"` // 用户唯一token key
	Status        string `gorm:"type:tinyint"`               // 状态 -1无效 1有效
	Level         string `gorm:"type:int"`                   // 等级
	Exp           string `gorm:"type:int"`                   // 经验
	AllowFileSize string `gorm:"type:char(255)"`             // 文件大小单位
	AllowQty      string `gorm:"type:bigint"`                // 允许储存的文件数
	AllowSuffix   string `gorm:"type:varchar(1024)"`         // 允许的文件后缀
}
