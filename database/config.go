package database


func def(key string) map[string]string {
	conf := make(map[string]string)
	conf["type"] = config.Get(key + ".type")
	conf["host"] = config.Get(key + ".host")
	conf["port"] = config.Get(key + ".port")
	conf["account"] = config.Get(key + ".account")
	conf["password"] = config.Get(key + ".password")
	conf["name"] = config.Get(key + ".name")
	conf["charset"] = config.Get(key + ".charset")
	return conf
}
