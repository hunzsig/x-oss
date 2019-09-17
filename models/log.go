package models

type Log struct {
	Id         string `gorm:"primary_key;type:bigint;not null;AUTO_INCREMENT"` // ID
	UserToken  string `gorm:"type:char(255)"`                                  // 用户token
	Msg        string `gorm:"type:varchar(1024)"`                              // 记录信息
	CreateTime string `gorm:"type:datetime"`                                   // 创建时间
}
