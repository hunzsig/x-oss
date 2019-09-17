package models

type Files struct {
	Hash         string `gorm:"primary_key;type:char(255);not null;unique"` // hash
	Name         string `gorm:"type:char(255)"`                             // 文件名
	TokenName    string `gorm:"type:char(255)"`                             // token名
	Suffix       string `gorm:"type:char(255)"`                             // 文件后缀
	Size         string `gorm:"type:bigint"`                                // 文件大小
	ContentType  string `gorm:"type:char(255)"`                             // 内容类型
	Path         string `gorm:"type:char(255)"`                             // 存在路径
	Uri          string `gorm:"type:char(255)"`                             // URI
	FromUrl      string `gorm:"type:char(255)"`                             // 来源地址
	CallQty      string `gorm:"type:bigint"`                                // 调用次数
	CallLastTime string `gorm:"type:datetime"`                              // 最后一次调用时间
	CreateTime   string `gorm:"type:datetime"`                              // 创建日期
	UpdateTime   string `gorm:"type:datetime"`                              // 更新日期
}
