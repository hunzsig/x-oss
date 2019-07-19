package main

import (
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

	oss := app.Party("/oss", func(ctx iris.Context) {
		token := ctx.Params().Get("token")
		if token != "abcdefg" {
			response.Error(ctx, "token break", nil)
		} else {
			ctx.Next()
		}
	})
	{
		// one file which is uploaded
		oss.Post("/upload/{type:string}", func(ctx iris.Context) {
			uploadType := ctx.Params().Get("type")
			response.Success(ctx, uploadType, nil)
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
