// +build enterprise

package migrations

import "dukeondope.ru/mlm/sandbox/config/admin"

func init() {
	AutoMigrate(&admin.QorMicroSite{})
}
