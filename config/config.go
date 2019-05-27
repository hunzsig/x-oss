package config

func DB(conf string) map[string]string {
	if conf == "default" {
		return mysql()
	}
	return nil
}
