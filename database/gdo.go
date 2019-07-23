package database

import (
	"../mapping"
	"../php2go"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

// DB 对象
type gdo struct {
	dbType  string
	dsn     string
	options map[string]string
	obj     *sql.DB
}

// 连接,并返回一个GDO
func GDO(name string) *gdo {
	link := get(name)
	dbDsn := dsn(link)
	switch link["type"] {
	case mapping.DBType.Mysql.Value:
	case mapping.DBType.Pgsql.Value:
	case mapping.DBType.Mssql.Value:
	case mapping.DBType.Sqlite.Value:
	default:
		panic("not support db type: " + link["type"])
	}
	dbObj, err := sql.Open(link["type"], dbDsn)
	if err != nil {
		panic(err)
	}
	gdo := new(gdo)
	gdo.dbType = link["type"]
	gdo.dsn = dbDsn
	gdo.obj = dbObj
	return gdo
}

func Mysql() *gdo {
	return GDO("mysql")
}

// query
func (gdo *gdo) Query(command string) (interface{}, error) {

	return true, nil
}

// 构建 select的sql 串
func (gdo *gdo) buildSelectSql() string {
	sqlStr := ""
	if gdo.obj == nil {
		panic("obj error")
	}
	if gdo.options["table"] == "" {
		panic("table error")
	}
	switch gdo.dbType {
	case mapping.DBType.Mysql.Value:
		sqlStr += "SELECT "
	case mapping.DBType.Pgsql.Value:
		sqlStr += "SELECT "
	default:
		panic("select tpl not support type:" + gdo.dbType)
	}
	return sqlStr
}

// 设置 Schema
func (gdo *gdo) Schema(val string) *gdo {
	if gdo.obj != nil {
		gdo.options["schema"] = val
	}
	return gdo
}

// 设置 table
func (gdo *gdo) Table(val string) *gdo {
	if gdo.obj != nil {
		gdo.options["table"] = val
	}
	return gdo
}

// 设置 field
func (gdo *gdo) Field(val string, table string) *gdo {
	if gdo.obj != nil {
		gdo.options["field"] = val
	}
	return gdo
}

// 设置 limit
func (gdo *gdo) Limit(val int) *gdo {
	if gdo.obj != nil {
		gdo.options["limit"] = strconv.Itoa(val)
	}
	return gdo
}

// 设置 offset
func (gdo *gdo) Offset(val int) *gdo {
	if gdo.obj != nil {
		gdo.options["offset"] = strconv.Itoa(val)
	}
	return gdo
}

// get one
func (gdo *gdo) One() map[string]string {
	result := make(map[string]string)
	gdo.options["limit"] = "1"
	itf, err := gdo.Query(gdo.buildSelectSql())
	if err != nil {
		return result
	}
	php2go.Dump(itf)
	return result
}

// 插入
func (gdo *gdo) Insert(data map[string]string) {

}
