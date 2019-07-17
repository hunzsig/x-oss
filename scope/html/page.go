package html

type Builder struct {
	Template string
	Params   map[string]string
}

func Build(tpl string, params map[string]string) *Builder {
	builder := new(Builder)
	builder.Template = tpl
	builder.Params = params
	return builder
}

/**
 * 构建器转换 html 内容
 */
func ToContent(builder *Builder) string {
	tpl := builder.Template
	params := builder.Params
	content := tpl
	return content
}

