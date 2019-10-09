package main

import (
	"./database"
	"./mapping"
	"./models"
	"./response"
	"./scope"
	"github.com/kataras/iris"
)

const fileMaxSize = 100 << 20      // 100MB
const multiFileMaxSize = 500 << 20 // 500MB

func route(app *iris.Application) {

	// default home
	app.Handle("GET", "/", func(ctx iris.Context) {
		scope.HomePage(ctx)
	})

	oss := app.Party("/oss/{token:string}", func(ctx iris.Context) {
		token := ctx.Params().Get("token")
		if token != "" {
			users := models.Users{}
			database.Mysql().Connect.Select([]string{"token", "status"}).Where("token = ?", token).First(&users)
			if users.Token == "" || users.Status != mapping.UserStatus.Enabled.Value {
				response.NotPermission(ctx, "token forbidden", nil)
				return
			}
			ctx.Params().Set("user_token", users.Token)
			ctx.Params().Set("user_exp", users.Exp)
			ctx.Next()
		} else {
			ctx.Next()
		}
	})
	{
		// one file which is uploaded
		oss.Post("/upload/{type:string}", func(ctx iris.Context) {
			/*
				todo 此处要判断文件大小/数量等
				token := ctx.Params().Get("token")
				users := models.Users{}
				database.Mysql().Connect.Select("token", "exp").Where("token = ?", token).First(&users)
			*/
			uploadType := ctx.Params().Get("type")
			if uploadType == "multi" {
				ctx.SetMaxRequestBodySize(multiFileMaxSize)
				scope.UploadMulti(ctx)
			} else {
				ctx.SetMaxRequestBodySize(fileMaxSize)
				scope.UploadOne(ctx)
			}
		})

		// download file by key
		oss.Get("/download/{fileKey:string}", func(ctx iris.Context) {
			scope.Download(ctx)
		})

		// get settings
		oss.Get("/settings", func(ctx iris.Context) {
			setting := scope.Settings()
			response.Success(ctx, "ok", setting)
		})

		// get files info
		oss.Get("/file", func(ctx iris.Context) {
			files := scope.FilesInfo(ctx)
			response.Success(ctx, "ok", files)
		})
		// get files list
		oss.Get("/files", func(ctx iris.Context) {
			files := scope.FilesList(ctx)
			response.Success(ctx, "ok", files)
		})
	}
}
