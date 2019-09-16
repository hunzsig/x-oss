package main

import (
	"./database"
	"./models"
	"./php2go"
	"./scope"
	"github.com/kataras/iris"
)

const fileMaxSize = 10 << 20      // 10MB
const multiFileMaxSize = 50 << 20 // 50MB

func route(app *iris.Application) {

	// default home
	app.Handle("GET", "/", func(ctx iris.Context) {
		scope.HomePage(ctx)
	})

	oss := app.Party("/oss/{token:string}", func(ctx iris.Context) {
		token := ctx.Params().Get("token")
		if token != "" {
			users := models.Users{}
			database.Mysql().Connect.LogMode(true)
			res := database.Mysql().Connect.Debug().First(&users)
			php2go.Dump(res)
			php2go.Dump(users)
			// response.NotPermission(ctx, "token forbidden", nil)
			ctx.Params().Set("user_token", users.Token)
			// ctx.Params().Set("user_exp", strconv.FormatInt(users.Exp))
		} else {
			ctx.Next()
		}
	})
	{
		// one file which is uploaded
		oss.Post("/upload/{type:string}", func(ctx iris.Context) {
			uploadType := ctx.Params().Get("type")
			if uploadType == "multi" {
				ctx.SetMaxRequestBodySize(multiFileMaxSize + 1<<20)
				scope.UploadMulti(ctx)
			} else {
				ctx.SetMaxRequestBodySize(fileMaxSize + 1<<20)
				scope.UploadOne(ctx)
			}
		})

		// download file by token
		oss.Get("/download/{fileKey:string}", func(ctx iris.Context) {
			scope.Download(ctx, ctx.Params().Get("fileKey"))
		})
	}
}
