package scope

import (
	"../database"
	"../models"
	"github.com/kataras/iris"
)

/**
 * 获取文件详情
 */
func FilesInfo(ctx iris.Context) models.Files {
	files := models.Files{}
	con := database.Mysql().Connect
	if ctx.Params().Get("hash") != "" {
		con.Where("`hash` like %?%", ctx.Params().Get("hash"))
	}
	if ctx.Params().Get("key") != "" {
		con.Where("`key` like %?%", ctx.Params().Get("key"))
	}
	if ctx.Params().Get("user_token") != "" {
		con.Where("`user_token` like %?%", ctx.Params().Get("user_token"))
	}
	if ctx.Params().Get("name") != "" {
		con.Where("`name` like %?%", ctx.Params().Get("name"))
	}
	con.First(&files)
	result := files
	return result
}

/**
 * 获取文件列表
 */
func FilesList(ctx iris.Context) []models.Files {
	var files []models.Files
	con := database.Mysql().Connect
	if ctx.Params().Get("hash") != "" {
		con.Where("`hash` like %?%", ctx.Params().Get("hash"))
	}
	if ctx.Params().Get("key") != "" {
		con.Where("`key` like %?%", ctx.Params().Get("key"))
	}
	if ctx.Params().Get("user_token") != "" {
		con.Where("`user_token` like %?%", ctx.Params().Get("user_token"))
	}
	if ctx.Params().Get("name") != "" {
		con.Where("`name` like %?%", ctx.Params().Get("name"))
	}
	con.Find(&files)
	result := files
	return result
}
