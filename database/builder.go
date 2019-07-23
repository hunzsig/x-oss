package database

import (
	"../mapping"
	"fmt"
)

// 构建 dsn
func dsn(conf map[string]string) string {
	dsnString := ""
	switch conf["type"] {
	case mapping.DBType.Mysql.Value:
		dsnString = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			conf["account"], conf["password"], conf["host"], conf["port"], conf["name"])
	case mapping.DBType.Pgsql.Value:
		dsnString = fmt.Sprintf(
			"port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf["port"], conf["account"], conf["password"], conf["name"])
	case mapping.DBType.Redis.Value:
		dsnString = fmt.Sprintf("%s:%s", conf["host"], conf["port"])
	}
	return dsnString
}

// 构建 select sql 模板
func selectTpl(conf map[string]string) string {
	tplStr := ""
	switch conf["type"] {
	case mapping.DBType.Mysql.Value:
		tplStr = ""
	case mapping.DBType.Pgsql.Value:
		tplStr = ""
	default:
		panic("select tpl not support type:" + conf["type"])
	}
	return tplStr
}
