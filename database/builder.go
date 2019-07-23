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
