// +build enterprise

package migrations

import (
	"github.com/qor/qor-example/app/admin"
)

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
