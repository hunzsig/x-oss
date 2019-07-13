package response

import (
	"github.com/kataras/iris"
)

func response(ctx iris.Context, code int16, msg string, data interface{}) {
	ctx.Gzip(true)
	ctx.JSON(iris.Map{"code": code, "msg": msg, "data": data})
}

func Success(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeSuccess, msg, data)
	return true
}

func Broadcast(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeBroadcast, msg, data)
	return true
}

func Goon(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeGoon, msg, data)
	return true
}

func Info(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeInfo, msg, data)
	return true
}

func Exception(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeException, msg, data)
	return false
}

func Error(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeError, msg, data)
	return false
}

func NotPermission(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeNotPermission, msg, data)
	return false
}

func NotFound(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeNotFound, msg, data)
	return false
}

func Abort(ctx iris.Context, msg string, data interface{}) bool {
	response(ctx, CodeAbort, msg, data)
	return false
}
