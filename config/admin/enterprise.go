// +build enterprise

package admin

import (
	"github.com/grengojbo/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/theplant/qor-enterprise/promotion"
)

func init() {
	// Benefits Definations
	discountRateArgument := Admin.NewResource(&struct {
		Percentage uint
	}{})
	discountRateArgument.Meta(&admin.Meta{Name: "Percentage", Label: "Percentage (e.g enter 10 for a 10% discount)"})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Rate",
		Resource: discountRateArgument,
	})

	discountAmountArgument := Admin.NewResource(&struct {
		Amount float32
	}{})
	discountAmountArgument.Meta(&admin.Meta{Name: "Amount", Label: "Amount (e.g enter 10 for a $10 discount)"})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Amount",
		Resource: discountAmountArgument,
	})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name: "Free Shipping",
	})

	// Rules Definations
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name: "Amount Greater Than",
		Resource: Admin.NewResource(&struct {
			Amount int
		}{}),
	})

	userGroupArgument := Admin.NewResource(&struct {
		Group string
	}{})
	userGroupArgument.Meta(&admin.Meta{Name: "Group", Type: "select_one", Collection: []string{"VIP", "Employee", "Normal"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "User Group",
		Resource: userGroupArgument,
	})

	// Auto migrations
	promotion.AutoMigrate(db.DB)

	// Add Promotions to Admin
	Admin.AddResource(&promotion.PromotionDiscount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}})
}
