package admin

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/qor-enterprise/microsite"
	"github.com/qor/admin"
	"github.com/qor/qor-example/config"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

func initMicrosite() {
	MicroSite = microsite.New(&microsite.Config{Dir: config.Root + "/public/microsites", Widgets: Widgets,
		URLProcessor: func(url string) string {
			reg := regexp.MustCompile(`/\w{2}-\w{2}/campaign`)
			if reg.MatchString(url) {
				return reg.ReplaceAllString(url, "/:locale/campaign")
			}
			return url
		},
	})
	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite"})
	Admin.AddResource(MicroSite)
	MicroSite.Funcs(func(http.ResponseWriter, *http.Request) template.FuncMap {
		return template.FuncMap{
			"say_hello":        func() string { return "Hello World" },
			"about_page_title": func() string { return "About Page Title" },
		}
	})
}
