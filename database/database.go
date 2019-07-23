package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/gomodule/redigo/redis"
)

// DB 对象
type gdo struct {
	dbType   string
	dsn      string
	objSqlDb *sql.DB
}

// 连接,并返回一个GDO
func Connect(name string) *gdo {
	link := get(name)
	dbDsn := dsn(link)
	dbObj, err := sql.Open(link["type"], dbDsn)
	if err != nil {
		panic(err)
	}
	gdo := new(gdo)
	gdo.dbType = link["type"]
	gdo.dsn = dbDsn
	gdo.objSqlDb = dbObj
	return gdo
}
func Mysql() *gdo {
	return Connect("mysql")
}
func Redis() *gdo {
	return Connect("redis")
}

// 插入
func (gdo *gdo) Table(data map[string]string) {
	gdo.db = db
}

// 插入
func (gdo *gdo) Insert(data map[string]string) {
	gdo.db = db
}
