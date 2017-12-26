// +build enterprise

package migrations

import "github.com/qor/qor-example/app/enterprise"

func init() {
	AutoMigrate(&enterprise.QorMicroSite{})
}
