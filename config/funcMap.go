package config

import (
	"github.com/qor/qor/i18n"
	"html/template"
)

var FuncMap = template.FuncMap{
	"renderInlineEditAssets": RenderInlineEditAssets,
	"t": T,
}

func T(key string, value string, args ...interface{}) template.HTML {
	return Config.I18n.EnableInlineEdit(true).Default(value).T("en-US", key, args)
}

func RenderInlineEditAssets() (template.HTML, error) {
	return i18n.RenderInlineEditAssets(true, true)
}
