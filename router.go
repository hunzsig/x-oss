package main

import (
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

	oss := app.Party("/oss", func(ctx iris.Context) {
		php2go.Dump(ctx)
		ctx.Next()
	})
	{
		// one file which is uploaded
		oss.Post("/upload/one", iris.LimitRequestBodySize(fileMaxSize), func(ctx iris.Context) {
			scope.UploadOne(ctx)
		})

		// multi files which is uploaded
		oss.Post("/upload/multi", iris.LimitRequestBodySize(multiFileMaxSize), func(ctx iris.Context) {
			scope.UploadMulti(ctx)
		})

		// download file by token
		oss.Get("/download/{token:string}", func(ctx iris.Context) {
			scope.Download(ctx)
		})
	}

}
