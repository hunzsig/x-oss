package database

import (
	"../mapping"
	"../php2go"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
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

// ---------------------------------------------

// 处理 key 为对应数据库形式
func (gdo *gdo) parseKey(key string) string {
	if key == "" {
		return key
	}
	key = php2go.Trim(key)
	if php2go.IsNumeric(key) {
		return key
	}
	match, err := regexp.MatchString(`[,'"*()`+"`"+`.\s]`, key)
	if err != nil || match == true {
		return key
	}
	switch gdo.dbType {
	case mapping.DBType.Mysql.Value:
		key = "`" + php2go.Trim(key) + "`"
	case mapping.DBType.Pgsql.Value:
		fallthrough
	case mapping.DBType.Mssql.Value:
		key = "\"" + php2go.Trim(key) + "\""
	case mapping.DBType.Sqlite.Value:
		key = "'" + php2go.Trim(key) + "'"
	}
	return key
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
		sqlStr += "select " + gdo.options["field"] + " from"
		if gdo.options["schema"] != "" {
			sqlStr += " " + gdo.options["schema"] + "." + gdo.options["table"]
		} else {
			sqlStr += " " + gdo.options["table"]
		}
		if gdo.options["join"] != "" {
			sqlStr += " " + gdo.options["join"]
		}
		if gdo.options["where"] != "" {
			sqlStr += " where " + gdo.options["where"]
		}
		if gdo.options["groupBy"] != "" {
			sqlStr += " group by " + gdo.options["groupBy"]
		}
		if gdo.options["orderBy"] != "" {
			sqlStr += " order by " + gdo.options["orderBy"]
		}
		if gdo.options["limit"] != "" {
			sqlStr += " " + gdo.options["limit"]
		}
	case mapping.DBType.Pgsql.Value:
		sqlStr += "select "
	default:
		panic("select tpl not support type:" + gdo.dbType)
	}
	return sqlStr
}

// 设置 Schema
func (gdo *gdo) Schema(val string) *gdo {
	if gdo.obj != nil {
		gdo.options["schema"] = gdo.parseKey(val)
	}
	return gdo
}

// 设置 table
func (gdo *gdo) Table(val string) *gdo {
	if gdo.obj != nil {
		gdo.options["table"] = gdo.parseKey(val)
	}
	return gdo
}

// 设置 field
// 以comma形式分割多个
func (gdo *gdo) Field(val string, table string) *gdo {
	if gdo.obj != nil {
		fieldArr := make([]string, 0)
		if gdo.options["field"] != "" {
			fieldArr = php2go.Explode(",", gdo.options["field"])
		}
		appendArr := php2go.Explode(",", val)
		for _, v := range appendArr {
			if !php2go.InArray(v, fieldArr) {
				fieldArr = append(fieldArr, gdo.parseKey(table)+gdo.parseKey(v))
			}
		}
		gdo.options["field"] = php2go.Implode(",", fieldArr)
	}
	php2go.Dump(gdo.options["field"])
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

// get multi
func (gdo *gdo) Multi() []map[string]string {
	result := make([]map[string]string, 0)
	itf, err := gdo.Query(gdo.buildSelectSql())
	if err != nil {
		return result
	}
	php2go.Dump(itf)
	return result
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
