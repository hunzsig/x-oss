package main

import (
	"./scope"
	"github.com/kataras/iris"
)

const fileMaxSize = 30 << 20 // 30MB

func route(app *iris.Application) {

	// default home
	app.Handle("ANY", "/", func(ctx iris.Context) {
		scope.Homepage(ctx)
	})

	app.Post("/upload", iris.LimitRequestBodySize(fileMaxSize+1<<20), func(ctx iris.Context) {
		scope.Upload(ctx)
	})

}
