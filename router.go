package main

import (
	"./database"
	"./response"
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
		if token != "abcdefg" {
			result, err := database.Mysql().Query("select * from `user` where `token` = " + token + " limit 1")
			if err != nil {
				response.Error(ctx, "token not allow", nil)
			}
			php2go.Dump(result)
			ctx.Params().Set("user_token",)
			ctx.Params().Set("user_exp",)
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
