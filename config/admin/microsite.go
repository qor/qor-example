// +build enterprise

package admin

import (
	"enterprise.getqor.com/microsite"
	"github.com/qor/admin"
	"github.com/qor/qor-example/config"
)

var MicroSite *microsite.MicroSite

type QorMicroSite struct {
	microsite.QorMicroSite
}

func init() {
	initWidgets()

	MicroSite = microsite.New(&microsite.Config{Dir: config.Root + "/public/microsites", Widgets: Widgets})
	MicroSite.Resource = Admin.AddResource(&QorMicroSite{}, &admin.Config{Name: "MicroSite"})

	Admin.AddResource(MicroSite)
}
