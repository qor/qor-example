// +build enterprise

package admin

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"enterprise.getqor.com/microsite"
	"github.com/qor/admin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/i18n"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

func init() {
	initWidgets()

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
	widgetsMeta := MicroSite.Resource.GetMeta("Widgets")
	widgetsMeta.SetFormattedValuer(func(site interface{}, ctx *qor.Context) interface{} {
		var results []template.HTML
		for _, widget := range site.(microsite.QorMicroSiteInterface).GetMicroSiteWidgets().Widgets {
			var setting QorWidgetSetting
			ctx.DB.First(&setting, "NAME = ?", widget.Name)
			results = append(results, template.HTML(`<img src="/images/Widget-`+setting.WidgetType+`.png" width="80" height="35" style="margin-right: 12px;"/><span>`+setting.Name+`</span>`))
		}
		return results
	})
	MicroSite.Funcs(func(site microsite.QorMicroSiteInterface, w http.ResponseWriter, req *http.Request) template.FuncMap {
		return template.FuncMap{
			"say_hello":        func() string { return "Hello World" },
			"about_page_title": func() string { return "About Page Title" },
			"t": func(key string, args ...interface{}) template.HTML {
				if len(args) == 0 {
					args = []interface{}{key}
				}
				key = fmt.Sprintf("microsite.%v.%v", site.GetMicroSiteID(), key)
				return inline_edit.InlineEdit(i18n.I18n, currentLocale(req), isEditMode(w, req))(key, args...)
			},
		}
	})
}

func currentLocale(req *http.Request) string {
	locale := "en-US"
	if cookie, err := req.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

func isEditMode(w http.ResponseWriter, req *http.Request) bool {
	return ActionBar.EditMode(w, req)
}
