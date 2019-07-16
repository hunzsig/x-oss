package main

import (
	"./response"
	"./system"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	system.Start()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)
	app.Use(recover.New())
	app.Use(logger.New())

	route(app)

	system.Dump(system.Env)

	app.Run(iris.Addr(":"+system.Env.Port), iris.WithoutServerError(iris.ErrServerClosed))

}

func notFound(ctx iris.Context) {
	response.Exception(ctx, "Not Found", nil)
}

func internalServerError(ctx iris.Context) {
	response.Exception(ctx, "Server Error", nil)
}
