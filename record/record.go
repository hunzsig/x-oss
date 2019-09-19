package record

import (
	"../database"
	"../models"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"os"
	"time"
)

/**
 * 数据库记录
 */
func Db(ctx iris.Context, message string) bool {
	token := ctx.Params().Get("user_token")
	if token == "" {
		token = "unknow"
	}
	log := models.Log{
		UserToken:  token,
		Msg:        message,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	database.Mysql().Connect.Create(&log)
	defer database.Mysql().Connect.Close()
	return true
}

/**
 * 文件记录
 */
func File(ctx iris.Context, message string) bool {
	fmt.Print(message)
	token := ctx.Params().Get("user_token")
	if token == "" {
		token = "unknow"
	}
	log := models.Log{
		UserToken: token,
		Msg:       message,
	}
	jsonBytes, _ := json.Marshal(log)
	path := "./logs/" + token + "/"
	os.MkdirAll(path, os.ModePerm)
	uri := path + time.Now().Format("20060102150405") + ".log"
	out, _ := os.OpenFile(uri, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()
	out.Write(jsonBytes)
	return true
}
