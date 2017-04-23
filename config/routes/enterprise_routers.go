// +build enterprise

package routes

import "github.com/qor/qor-example/config/admin"

func init() {
	Router()
	WildcardRouter.AddHandler(admin.MicroSite)
}
