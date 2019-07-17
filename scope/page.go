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
	tpl404    = `<div>{{error}}</div>`
)

func page(ctx iris.Context, page string) bool {
	htmlBuilder := new(html.HtmlBuilder)
	switch page {
	case "home":
		htmlBuilder.Template, pageError = php2go.FileGetContents("html/home.html")
		htmlBuilder.Params = make(map[string]string)
		htmlBuilder.Params["title"] = "h-assets"
		htmlBuilder.Params["tips"] = "Welcome!~h-assets"
	default:
		htmlBuilder.Template = tpl404
		htmlBuilder.Params = make(map[string]string)
		htmlBuilder.Params["error"] = "page #" + page + " not found"
	}
	if pageError != nil {
		htmlBuilder.Template = tplError
		htmlBuilder.Params = make(map[string]string)
		htmlBuilder.Params["error"] = pageError.Error()
	}
	php2go.Dump(htmlBuilder)
	return response.Html(ctx, html.ToContent(htmlBuilder))
}

func HomePage(ctx iris.Context) bool {
	return page(ctx, "home")
}
