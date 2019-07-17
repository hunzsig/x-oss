package html

import (
	"../php2go"
)

type HtmlBuilder struct {
	Template string
	Params   map[string]string
}

func Build(tpl string, params map[string]string) *HtmlBuilder {
	builder := new(HtmlBuilder)
	builder.Template = tpl
	builder.Params = params
	return builder
}

/**
 * 构建器转换 html 内容
 */
func ToContent(builder *HtmlBuilder) string {
	tpl := builder.Template
	params := builder.Params
	content := tpl
	for k, v := range params {
		tplStr := "{{" + k + "}}"
		content = php2go.StrReplace(tplStr, v, content, -1)
	}
	return content
}
