// +build enterprise

package migrations

import (
	"enterprise.getqor.com/microsite"
	"github.com/qor/qor-example/config/admin"
)

func init() {
	AutoMigrate(&admin.QorMicroSite{}, &microsite.QorMicorSiteWidgetSetting{})
}
