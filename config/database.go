package config

func mysql() map[string]string {
	conf := make(map[string]string)
	conf["type"] = "mysql"
	conf["host"] = "127.0.0.1"
	conf["port"] = "3306"
	conf["account"] = "root"
	conf["password"] = "123456"
	conf["name"] = "workflow"
	conf["charset"] = "utf8mb4"
	return conf
}
