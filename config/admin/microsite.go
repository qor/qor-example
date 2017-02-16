// +build enterprise

package admin

import (
	"regexp"

	"enterprise.getqor.com/microsite"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor-example/config"
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
	})

	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite"})

	Admin.AddResource(MicroSite)
}
