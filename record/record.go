package record

import (
	"../database"
	"../models"
	"github.com/kataras/iris"
)

/**
 * 数据库记录
 */
func Db(ctx iris.Context, message string) bool {
	token := ctx.Params().Get("user_token")
	if token == "" {
		token = ctx.Params().Get("token")
	}
	log := models.Log{
		UserToken: token,
		Msg:       message,
	}
	database.Mysql().Connect.NewRecord(&log)
	defer database.Mysql().Connect.Close()
	return true
}
