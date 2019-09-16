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
	mysql["charset"] = env.Data.MysqlCharset
	set("mysql", mysql)

	// pgsql
	pgsql := make(map[string]string)
	pgsql["type"] = "pgsql"
	pgsql["host"] = env.Data.PgsqlHost
	pgsql["port"] = env.Data.PgsqlPort
	pgsql["account"] = env.Data.PgsqlAccount
	pgsql["password"] = env.Data.PgsqlPassword
	pgsql["name"] = env.Data.PgsqlName
	pgsql["charset"] = env.Data.PgsqlCharset
	set("pgsql", pgsql)

	// mssql
	mssql := make(map[string]string)
	mssql["type"] = "mssql"
	mssql["host"] = env.Data.MssqlHost
	mssql["port"] = env.Data.MssqlPort
	mssql["account"] = env.Data.MssqlAccount
	mssql["password"] = env.Data.MssqlPassword
	mssql["name"] = env.Data.MssqlName
	mssql["charset"] = env.Data.MssqlCharset
	set("mssql", mssql)

	// sqlite
	sqlite := make(map[string]string)
	sqlite["type"] = "sqlite"
	sqlite["path"] = env.Data.SqlitePath
	sqlite["charset"] = env.Data.SqliteCharset
	set("sqlite", sqlite)

	// redis
	redis := make(map[string]string)
	redis["type"] = "redis"
	redis["host"] = env.Data.RedisHost
	redis["port"] = env.Data.RedisPort
	redis["password"] = env.Data.RedisPassword
	set("redis", redis)
}
