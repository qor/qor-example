package config

import (
	"html/template"
)

var FuncMap = template.FuncMap{
	"t": T,
}

func T(key string, value string, args ...interface{}) template.HTML {
	return Config.I18n.Default(value).T("en-US", key, args)
}
