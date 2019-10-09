package scope

import (
	"../database"
	"../models"
	"../php2go"
)

type SettingResult struct {
	Key  string
	Data []string
}

/**
 * 获取settings
 */
func Settings() []SettingResult {
	var setting []models.Setting
	con := database.Mysql().Connect
	con.Find(&setting)
	var result []SettingResult
	for _, v := range setting {
		data := php2go.Explode(",", v.Data)
		result = append(result, SettingResult{
			v.Key,
			data,
		})
	}
	return result
}
