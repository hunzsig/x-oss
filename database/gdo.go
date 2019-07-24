package database

import (
	"../mapping"
	"../php2go"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/core/errors"
	"regexp"
	"strconv"
	"strings"
)

// DB 对象
type gdo struct {
	dbType  string
	dsn     string
	obj     *sql.DB
	options map[string]string
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
	gdo.options = make(map[string]string)
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
	php2go.Dump(key)
	key = php2go.Trim(key)
	if php2go.IsNumeric(key) {
		return key
	}
	php2go.Dump(key)
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
func (gdo *gdo) Query(query string) (interface{}, error) {
	php2go.Dump(gdo.options)
	php2go.Dump(query)
	queryItems := php2go.Explode(" ", query)
	if gdo.options["table"] == "" {
		return false, errors.New("lose table")
	}
	statement := strings.ToLower(queryItems[0])
	//read model,check cache
	if statement == "select" || statement == "show" {
		// cache
	}
	// 执行新一轮的查询，并释放上一轮结果
	if statement == "select" || statement == "show" {
		rows, err := gdo.obj.Query(query)
		if err != nil {
			return false, err
		}
		php2go.Dump(rows)
	} else if statement == "insert" {
		result, err := gdo.obj.Exec(query)
		if err != nil {
			return false, err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return false, err
		}
		return id, nil
	} else if statement == "update" || statement == "delete" {
		result, err := gdo.obj.Exec(query)
		if err != nil {
			return false, err
		}
		num, err := result.RowsAffected()
		if err != nil {
			return false, err
		}
		return num, nil
	}
	return true, nil
}

// 构建 select的sql 串
func (gdo *gdo) buildSelectSql() string {
	sqlStr := ""
	if gdo.obj == nil {
		return sqlStr
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
			sqlStr += " limit" + gdo.options["limit"]
		}
		if gdo.options["offset"] != "" {
			sqlStr += " offset" + gdo.options["offset"]
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
				fieldArr = append(fieldArr, gdo.parseKey(table)+"."+gdo.parseKey(v))
			}
		}
		gdo.options["field"] = php2go.Implode(",", fieldArr)
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
	gdo.Limit(1)
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
