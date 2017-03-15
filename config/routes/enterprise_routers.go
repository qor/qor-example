// +build enterprise

package routes

import "dukeondope.ru/mlm/sandbox/config/admin"

func init() {
	Router()
	WildcardRouter.AddHandler(admin.MicroSite)
}
