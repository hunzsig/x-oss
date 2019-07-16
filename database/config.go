package database

import (
	"../system"
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
	mysql["host"] = system.Env.MysqlHost
	mysql["port"] = system.Env.MysqlPort
	mysql["account"] = system.Env.MysqlAccount
	mysql["password"] = system.Env.MysqlPassword
	mysql["name"] = system.Env.MysqlName
	set("mysql", mysql)
	// redis
	redis := make(map[string]string)
	redis["type"] = "redis"
	redis["host"] = system.Env.RedisHost
	redis["port"] = system.Env.RedisPort
	redis["password"] = system.Env.RedisPassword
	set("redis", redis)
}
