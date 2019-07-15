package system

import "time"

type envReflect struct {
	App     string
	Port    int      `default:"8080"`
	Debug   bool     `default:"0"`
	Hosts   []string `slice_sep:","`
	Timeout time.Duration

	RedisHost string `default:"127.0.0.1"`
	RedisPort int    `default:"6379"`

	MysqlHost string `default:"127.0.0.1"`
	MysqlPort int    `default:"3306"`
}