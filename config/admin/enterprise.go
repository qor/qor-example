// +build enterprise

package admin

import (
	"fmt"

	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/admin"
	"github.com/theplant/qor-enterprise/promotion"
)

func init() {
	// Benefits Definations
	discountRateArgument := Admin.NewResource(&struct {
		Percentage uint
		Category   string
	}{})
	discountRateArgument.Meta(&admin.Meta{Name: "Percentage", Label: "Percentage (e.g enter 10 for a 10% discount)"})
	discountRateArgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Rate",
		Resource: discountRateArgument,
	})

	discountAmountArgument := Admin.NewResource(&struct {
		Amount   float32
		Category string
	}{})
	discountAmountArgument.Meta(&admin.Meta{Name: "Amount", Label: "Amount (e.g enter 10 for a $10 discount)"})
	discountAmountArgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Discount Amount",
		Resource: discountAmountArgument,
	})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name: "shipping fee",
		Resource: Admin.NewResource(&struct {
			Price float32
		}{}),
	})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name: "2th day shipping fee",
		Resource: Admin.NewResource(&struct {
			Price float32
		}{}),
	})

	productCodeCollection := func(value interface{}, context *qor.Context) [][]string {
		var products []models.Product
		var results [][]string
		context.GetDB().Find(&products)
		for _, product := range products {
			results = append(results, []string{fmt.Sprint(product.ID), product.Code})
		}
		return results
	}
	productWithPriceArgument := Admin.NewResource(&struct {
		ProductCode string
		Quantity    uint
		Price       float32
	}{})
	productWithPriceArgument.Meta(&admin.Meta{Name: "ProductCode", Type: "select_one", Collection: productCodeCollection})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Product With Price",
		Resource: productWithPriceArgument,
	})

	// Rules Definations
	amountGreaterThanArgument := Admin.NewResource(&struct {
		Amount   int
		Category string
	}{})
	amountGreaterThanArgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Amount Greater Than",
		Resource: amountGreaterThanArgument,
	})

	userGroupArgument := Admin.NewResource(&struct {
		Group string
	}{})
	userGroupArgument.Meta(&admin.Meta{Name: "Group", Type: "select_one", Collection: []string{"VIP", "Employee", "Normal"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "User Group",
		Resource: userGroupArgument,
	})

	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name: "From Link",
		Resource: Admin.NewResource(&struct {
			URL string
		}{}),
	})

	hasProductrgument := Admin.NewResource(&struct {
		ProductCode string
		Category    string
	}{})
	hasProductrgument.Meta(&admin.Meta{Name: "ProductCode", Type: "select_one", Collection: productCodeCollection})
	hasProductrgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Has Product",
		Resource: hasProductrgument,
	})

	// Auto migrations
	promotion.AutoMigrate(db.DB)

	// Add Promotions to Admin
	Admin.AddResource(&promotion.PromotionDiscount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}})
}
