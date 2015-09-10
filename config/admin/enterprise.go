// +build enterprise

package admin

import (
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/theplant/qor-enterprise/promotion"
)

func init() {
	// Benefits Definations
	type discountRateArgument struct {
		Percentage uint
	}
	discountRateArgumentResource := Admin.NewResource(&discountRateArgument{})
	discountRateArgumentResource.Meta(&admin.Meta{Name: "Percentage", Label: "Percentage (e.g enter 10 for a 10% discount)"})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Rate",
		Resource: discountRateArgumentResource,
	})

	type discountAmountArgument struct {
		Amount float32
	}
	discountAmountArgumentResource := Admin.NewResource(&discountAmountArgument{})
	discountAmountArgumentResource.Meta(&admin.Meta{Name: "Amount", Label: "Amount (e.g enter 10 for a 10$ discount)"})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Amount",
		Resource: discountAmountArgumentResource,
	})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name: "Free Shipping",
	})

	// Rules Definations
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

	type hasProductArgument struct {
		ProductCode string
	}
	hasProductResource := Admin.NewResource(&hasProductArgument{})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Has Product",
		Resource: hasProductResource,
	})

	promotion.AutoMigrate(db.DB)
	Admin.AddResource(&promotion.PromotionDiscount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}})
}
