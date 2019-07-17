package scope

import (
	"../response"
	"github.com/kataras/iris"
)

func Page(ctx iris.Context) bool {
	return response.Html(ctx, "Welcome!~h-assets", nil)
}
