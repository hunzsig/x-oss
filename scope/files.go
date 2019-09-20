package scope

import (
	"../database"
	"../models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

func filesWhere(ctx iris.Context, con *gorm.DB) *gorm.DB {
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
	return con
}

/**
 * 获取文件详情
 */
func FilesInfo(ctx iris.Context) models.Files {
	files := models.Files{}
	con := database.Mysql().Connect
	con = filesWhere(ctx, con)
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
	con = filesWhere(ctx, con)
	con.Find(&files)
	result := files
	return result
}
