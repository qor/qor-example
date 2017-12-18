// +build enterprise

package migrations

import (
	"github.com/qor/qor-example/config/admin"
)

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
