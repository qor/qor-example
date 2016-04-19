package i18n

import (
	"fmt"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/i18n/backends/yaml"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	"html/template"
	"path/filepath"
)

var I18n *i18n.I18n

var FuncMap = template.FuncMap{
	"t": T,
}

func init() {
	fmt.Printf("aaa-------------------%v\n", db.DB)
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
	for key, value := range FuncMap {
		config.View.RegisterFuncMap(key, value)
	}
}

func T(values ...interface{}) template.HTML {
	switch len(values) {
	case 1:
		return I18n.EnableInlineEdit(true).T("en-US", fmt.Sprint(values[0]))
	case 2:
		return I18n.EnableInlineEdit(true).T("en-US", fmt.Sprint(values[0]), values[1:]...)
	default:
		panic("passed wrong params for T")
	}
	return ""
}
