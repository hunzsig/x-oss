package env

import "time"

type envReflect struct {
	App     string
	Port    string   `default:"8080"`
	Debug   bool     `default:"0"`
	Hosts   []string `default:"127.0.0.1"`
	Timeout time.Duration

	RedisHost     string `default:"127.0.0.1"`
	RedisPort     string `default:"6379"`
	RedisPassword string `default:""`

	MysqlHost     string `default:"127.0.0.1"`
	MysqlPort     string `default:"3306"`
	MysqlAccount  string `default:"root"`
	MysqlPassword string `default:"123456"`
	MysqlName     string `default:"hunzsig"`
	MysqlCharset  string `default:"utf8mb4"`

	PgsqlHost     string `default:"127.0.0.1"`
	PgsqlPort     string `default:"5432"`
	PgsqlAccount  string `default:"root"`
	PgsqlPassword string `default:"123456"`
	PgsqlName     string `default:"hunzsig"`
	PgsqlCharset  string `default:"utf8"`

	MssqlHost     string `default:"127.0.0.1"`
	MssqlPort     string `default:"1433"`
	MssqlAccount  string `default:"sa"`
	MssqlPassword string `default:"123456"`
	MssqlName     string `default:"hunzsig"`
	MssqlCharset  string `default:"utf8"`

	SqlitePath    string `default:"/tmp/hunzsig.db"`
}
