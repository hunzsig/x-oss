package scope

import (
	"../response"
	"github.com/kataras/iris"
)

func Homepage(ctx iris.Context) bool {
	return response.Success(ctx, "Welcome!~h-assets", nil)
}
