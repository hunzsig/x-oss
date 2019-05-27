package main

import (
	"./system"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	system.Start()

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	route(app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))

}
