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
			"%s:%s@tcp(%s:%s)/%s?parseTime=True&charset=%s",
			conf["account"], conf["password"], conf["host"], conf["port"], conf["name"], conf["charset"])
	case mapping.DBType.Pgsql.Value:
		dsnString = fmt.Sprintf(
			"host=%s:%s user=%s password=%s dbname=%s charset=%s sslmode=disable",
			conf["host"], conf["port"], conf["account"], conf["password"], conf["name"], conf["charset"])
	case mapping.DBType.Redis.Value:
		dsnString = fmt.Sprintf("%s:%s", conf["host"], conf["port"])
	}
	return dsnString
}
