package response

import (
	"github.com/kataras/iris"
)

func Html(ctx iris.Context, htmlcontent string) bool {
	ctx.HTML(htmlcontent)
	return true
}

func Download(ctx iris.Context, filename string) bool {
	ctx.SendFile(filename, "helloworld.log")
	return true
}

func Json(ctx iris.Context, code int16, msg string, data interface{}) {
	ctx.Gzip(true)
	ctx.JSON(iris.Map{"code": code, "msg": msg, "data": data})
}