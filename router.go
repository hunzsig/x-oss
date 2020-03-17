package main

import (
	"github.com/kataras/iris"
	"x-oss/database"
	"x-oss/env"
	"x-oss/mapping"
	"x-oss/models"
	"x-oss/php2go"
	"x-oss/response"
	"x-oss/scope"
)

const fileMaxSize = 100 << 20      // 100MB
const multiFileMaxSize = 500 << 20 // 500MB

func enter(ctx iris.Context) {
	if php2go.InArray(ctx.RemoteAddr(), env.Data.Hosts) == false {
		response.NotFound(ctx, "x-oss", nil)
	} else {
		ctx.Next()
	}
}

func route(app *iris.Application) {

	// before enter
	// 根据env的Hosts配置，禁止除此IP的服务器访问
	app.Use(enter)

	// default home
	app.Handle("GET", "/", func(ctx iris.Context) {
		scope.HomePage(ctx)
	})

	// get settings
	app.Get("/settings", func(ctx iris.Context) {
		setting := scope.Settings()
		response.Success(ctx, "ok", setting)
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
