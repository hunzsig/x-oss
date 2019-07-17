package main

import (
	"./scope"
	"github.com/kataras/iris"
)

const fileMaxSize = 30 << 20 // 30MB

func route(app *iris.Application) {

	// default home
	app.Handle("GET", "/", func(ctx iris.Context) {
		scope.HomePage(ctx)
	})

	oss := app.Party("/oss")
	{
		// handle files which is uploaded
		oss.Post("/upload", iris.LimitRequestBodySize(fileMaxSize+1<<20), func(ctx iris.Context) {
			scope.Upload(ctx)
		})

		// download file by token
		oss.Get("/download/{token:string}", func(ctx iris.Context) {
			scope.Download(ctx)
		})
	}

}
