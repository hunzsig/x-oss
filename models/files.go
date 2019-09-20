package models

type Files struct {
	Hash         string `gorm:"primary_key;type:char(255);not null;unique"` // hash（这个哈希是根据文件二进制数据合成的）
	Key          string `gorm:"type:char(255);not null;unique"`             // 文件key（这个key是生成的，用于访问）
	UserToken    string `gorm:"type:char(255)"`                             // 用户token
	Name         string `gorm:"type:char(255)"`                             // 文件名
	Md5Name      string `gorm:"type:char(255)"`                             // md5名字（这是合成的md5名字）
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
