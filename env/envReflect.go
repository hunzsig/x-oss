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
}
