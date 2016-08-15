package admin

import (
	"html/template"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/microsite"
	"github.com/qor/qor-example/config"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

func init() {
	MicroSite = microsite.New(config.Root+"/public/microsites", Widgets)
	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite"})
	Admin.AddResource(MicroSite)
	MicroSite.Funcs(func(http.ResponseWriter, *http.Request) template.FuncMap {
		return template.FuncMap{
			"say_hello":        func() string { return "Hello World" },
			"about_page_title": func() string { return "About Page Title" },
		}
	})
}
