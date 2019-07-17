package database

import (
	"../env"
)

var (
	conf map[string]map[string]string
)

func set(name string, setting map[string]string) {
	conf[name] = setting
}

func get(name string) map[string]string {
	return conf[name]
}

func init() {
	// conf
	conf = make(map[string]map[string]string)
	// mysql
	mysql := make(map[string]string)
	mysql["type"] = "mysql"
	mysql["host"] = env.Data.MysqlHost
	mysql["port"] = env.Data.MysqlPort
	mysql["account"] = env.Data.MysqlAccount
	mysql["password"] = env.Data.MysqlPassword
	mysql["name"] = env.Data.MysqlName
	set("mysql", mysql)
	// redis
	redis := make(map[string]string)
	redis["type"] = "redis"
	redis["host"] = env.Data.RedisHost
	redis["port"] = env.Data.RedisPort
	redis["password"] = env.Data.RedisPassword
	set("redis", redis)
}
