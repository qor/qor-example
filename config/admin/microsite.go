// +build enterprise

package admin

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path"
	"regexp"
	"strings"

	"enterprise.getqor.com/microsite"
	"github.com/qor/admin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor/utils"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

func init() {
	initWidgets()

	MicroSite = microsite.New(&microsite.Config{Dir: config.Root + "/public/microsites", Widgets: Widgets,
		URLProcessor: func(url string, context *qor.Context) string {
			reg := regexp.MustCompile(`/\w{2}-\w{2}/campaign`)
			if reg.MatchString(url) {
				return reg.ReplaceAllString(url, "/:locale/campaign")
			}
			return url
		},
		TemplateFinder: func(url string, site microsite.QorMicroSiteInterface, context *qor.Context) (io.ReadSeeker, error) {
			reg := regexp.MustCompile(`/:locale/campaign/code`)
			if reg.MatchString(url) {
				return strings.NewReader("Campaign Pomotion code: AH0134"), nil
			}

			reg = regexp.MustCompile(`/:locale/campaign/blogs/.+`)
			if reg.MatchString(url) {
				pak := site.GetCurrentPackage()
				return pak.GetTemplate(MicroSite, site, "/:locale/campaign/blogs/show.html")
			}
			return nil, microsite.ErrNotFound
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
			results = append(results, template.HTML(`<img src="/images/Widget-`+setting.WidgetType+`.png" /><span>`+setting.Name+`</span>`))
		}
		return results
	})
	MicroSite.Funcs(func(site microsite.QorMicroSiteInterface, w http.ResponseWriter, req *http.Request) template.FuncMap {
		return template.FuncMap{
			"say_hello":        func() string { return "Hello World" },
			"about_page_title": func() string { return "About Page Title" },
			"blog_title": func() string {
				url := req.URL.Path
				reg := regexp.MustCompile(`/\w{2}-\w{2}/campaign/blogs/.+`)
				if reg.MatchString(url) {
					return utils.HumanizeString(path.Base(url))
				}
				return ""
			},
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
