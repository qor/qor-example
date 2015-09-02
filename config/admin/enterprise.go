// +build enterprise

package admin

import (
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/theplant/qor-enterprise/promotion"
)

func init() {
	type discountRateArgument struct {
		Value uint
	}
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Rate",
		Resource: Admin.NewResource(&discountRateArgument{}),
	})

	type discountValueArgument struct {
		Value float32
	}
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Value",
		Resource: Admin.NewResource(&discountValueArgument{}),
	})

	type amountArgument struct {
		Amount int
	}
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Amount Greater Than",
		Resource: Admin.NewResource(&amountArgument{}),
	})

	type userGroupArgument struct {
		Group string
	}
	userGroupResource := Admin.NewResource(&userGroupArgument{})
	userGroupResource.Meta(&admin.Meta{Name: "Group", Type: "select_one", Collection: []string{"VIP", "Employee", "Normal"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "User Group",
		Resource: userGroupResource,
	})

	promotion.AutoMigrate(db.DB)
	Admin.AddResource(&promotion.PromotionDiscount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}})
}
