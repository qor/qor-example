// +build enterprise

package admin

import (
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/theplant/qor-enterprise/promotion"
)

func init() {
	promotion.AutoMigrate(db.DB)
	Admin.AddResource(&promotion.PromotionDiscount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}})
}
