package scope

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"x-oss/database"
	"x-oss/models"
)

func filesWhere(ctx iris.Context, con *gorm.DB) *gorm.DB {
	con = con.Where("`user_token` = ?", ctx.Params().Get("user_token"))
	if ctx.FormValue("hash") != "" {
		con = con.Where("`hash` like ?", "%"+ctx.FormValue("hash")+"%")
	}
	if ctx.FormValue("key") != "" {
		con = con.Where("`key` like ?", "%"+ctx.FormValue("key")+"%")
	}
	if ctx.FormValue("user_token") != "" {
		con = con.Where("`user_token` like ?", "%"+ctx.FormValue("user_token")+"%")
	}
	if ctx.FormValue("name") != "" {
		con = con.Where("`name` like ?", "%"+ctx.FormValue("name")+"%")
	}
	return con
}

/**
 * 获取文件详情
 */
func FilesInfo(ctx iris.Context) models.Files {
	files := models.Files{}
	con := database.Mysql().Connect.Debug()
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
