package response

import (
	"../record"
	"github.com/kataras/iris"
)

func rec(ctx iris.Context, code int16, msg string) {
	if code == CodeException || code == CodeError || code == CodeNotPermission || code == CodeNotFound || code == CodeAbort {
		record.File(ctx, msg)
		record.Db(ctx, msg)
	}
}

func Html(ctx iris.Context, htmlcontent string) bool {
	ctx.Gzip(true)
	ctx.HTML(htmlcontent)
	return true
}

func Download(ctx iris.Context, filename string, destinationName string) bool {
	ctx.SendFile(filename, destinationName)
	return true
}

func Json(ctx iris.Context, code int16, msg string, data interface{}) {
	rec(ctx, code, msg)
	ctx.Gzip(true)
	ctx.JSON(iris.Map{"code": code, "msg": msg, "data": data})
}
