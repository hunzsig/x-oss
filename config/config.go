package config

var stack map[string]string

/**
 * 设定配置值
 */
func Set(key string, val string) {
	stack[key] = val
}

/**
 * 获取配置值
 */
func Get(key string) string {
	return stack[key]
}
