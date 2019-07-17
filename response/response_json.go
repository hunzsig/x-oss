package response

import (
	"github.com/kataras/iris"
)

func Success(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeSuccess, msg, data)
	return true
}

func Broadcast(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeBroadcast, msg, data)
	return true
}

func Goon(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeGoon, msg, data)
	return true
}

func Info(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeInfo, msg, data)
	return true
}

func Exception(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeException, msg, data)
	return false
}

func Error(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeError, msg, data)
	return false
}

func NotPermission(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeNotPermission, msg, data)
	return false
}

func NotFound(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeNotFound, msg, data)
	return false
}

func Abort(ctx iris.Context, msg string, data interface{}) bool {
	Json(ctx, CodeAbort, msg, data)
	return false
}
