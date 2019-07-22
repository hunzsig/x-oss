package database

/**
 * GDO 对象
 */
type gdo struct {
	Dsn func()
}

/**
 * 初始化一个GDO
 */
func GDO() *gdo {
	obj := new(gdo)
	return obj
}
