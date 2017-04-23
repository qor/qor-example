// +build enterprise

package admin

import (
	"fmt"

	"enterprise.getqor.com/promotion"
	"github.com/qor/admin"
	"github.com/qor/qor"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
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
		Name: "Shipping Fee",
		Resource: Admin.NewResource(&struct {
			Price float32
		}{}),
	})

	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name: "2nd Day Shipping Fee",
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
	combinedDiscountArgument := Admin.NewResource(&struct {
		ProductCodes []string
		Category     string
		Quantity     uint
		Price        float32
		Percentage   uint
		Discount     uint
	}{})
	combinedDiscountArgument.Meta(&admin.Meta{Name: "ProductCodes", Type: "select_many", Collection: productCodeCollection})
	combinedDiscountArgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	combinedDiscountArgument.Meta(&admin.Meta{Name: "Percentage", Label: "Discount Percentage (e.g enter 10 for a 10% discount)"})
	combinedDiscountArgument.Meta(&admin.Meta{Name: "Discount", Label: "Discount Amount (e.g enter 10 for a $10 discount)"})
	promotion.RegisterBenefitHandler(promotion.BenefitHandler{
		Name:     "Combined Discounts",
		Resource: combinedDiscountArgument,
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

	quantityGreaterThanArgument := Admin.NewResource(&struct {
		ProductCodes []string
		Category     string
		Quantity     int
	}{})
	quantityGreaterThanArgument.Meta(&admin.Meta{Name: "ProductCodes", Type: "select_many", Collection: productCodeCollection})
	quantityGreaterThanArgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Quantity Greater Than",
		Resource: quantityGreaterThanArgument,
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
			VariableName string
			Value        string
		}{}),
	})

	hasProductrgument := Admin.NewResource(&struct {
		ProductCodes []string
		Category     string
	}{})
	hasProductrgument.Meta(&admin.Meta{Name: "ProductCodes", Type: "select_many", Collection: productCodeCollection})
	hasProductrgument.Meta(&admin.Meta{Name: "Category", Type: "select_one", Collection: []string{"All Products", "Bags", "Summer Shirts", "Pants"}})
	promotion.RegisterRuleHandler(promotion.RuleHandler{
		Name:     "Has Product",
		Resource: hasProductrgument,
	})

	// Auto migrations
	promotion.AutoMigrate(db.DB)

	// Add Promotions to Admin
	Admin.AddResource(&promotion.Discount{}, &admin.Config{Name: "Promotions", Menu: []string{"Site Management"}, Priority: 3})
}
