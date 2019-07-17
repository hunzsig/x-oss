package scope

import (
	"../html"
	"../php2go"
	"../response"
	"github.com/kataras/iris"
)

var (
	pageError error
	tplError  = `<div>{{error}}</div>`
)

func page(ctx iris.Context, page string, params map[string]string) bool {
	htmlBuilder := new(html.HtmlBuilder)
	htmlBuilder.Template, pageError = php2go.FileGetContents("html/" + page + ".html")
	htmlBuilder.Params = params
	if pageError != nil {
		htmlBuilder.Template = tplError
		htmlBuilder.Params = make(map[string]string)
		htmlBuilder.Params["error"] = pageError.Error()
	}
	return response.Html(ctx, html.ToContent(htmlBuilder))
}

func HomePage(ctx iris.Context) bool {
	params := make(map[string]string)
	params["title"] = "h-assets"
	params["tips"] = "Welcome!~h-assets"
	return page(ctx, "home", params)
}
