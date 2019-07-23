package database

import (
	"fmt"
)

func dsn(conf map[string]string) string {
	dsnString := ""
	switch conf["type"] {
	case "mysql":
		dsnString = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			conf["account"], conf["password"], conf["host"], conf["port"], conf["name"])
	case "pysql":
		dsnString = fmt.Sprintf(
			"port=$s user=%s password=%s dbname=%s sslmode=disable", "postgres", "123456", "postgres")
	case "redis":
		dsnString = fmt.Sprintf("%s:%s", conf["host"], conf["port"])
	}
	return dsnString
}
