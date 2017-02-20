package controllers

import (
	"fmt"

	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/seo"
	qor_seo "github.com/qor/seo"
	// "github.com/qor/transition"
)

func ProductShow(ctx *gin.Context) {
	var (
		product        models.Product
		colorVariation models.ColorVariation
		codes          = strings.Split(ctx.Param("code"), "_")
		productCode    = codes[0]
		colorCode      string
	)

	if len(codes) > 1 {
		colorCode = codes[1]
	}

	if DB(ctx).Where(&models.Product{Code: productCode}).First(&product).RecordNotFound() {
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusFound)
	}

	DB(ctx).Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"product_show",
		gin.H{
			"ActionBarTag":   admin.ActionBar.Actions(action_bar.EditResourceAction{Value: product, Inline: true, EditModeOnly: true}).Render(ctx.Writer, ctx.Request),
			"Product":        product,
			"ColorVariation": colorVariation,
			"SEOTag":         seo.SEOCollection.Render(&qor.Context{DB: DB(ctx)}, "Product Page", product),
			"MicroProduct": qor_seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
				Image:       colorVariation.MainImageURL(),
			}.Render(),
			"Categories":    CategoriesList(ctx),
			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func AddToCart(ctx *gin.Context) {
	var (
		product     models.Product
		user        = CurrentUser(ctx)
		codes       = strings.Split(ctx.PostForm("code"), "_")
		productCode = codes[0]
		order       models.Order
		// orderItems  []models.OrderItem
		// OrderStateMachine = transition.New(&order)
	)
	DB(ctx).Where(&models.Product{Code: productCode}).First(&product)

	if err := DB(ctx).Create(&order).Error; err != nil {
		fmt.Printf("create order (%v) failure, got err %v", order, err)
	}

	order.UserID = user.ID
	// order.ShippingAddressID = user.Addresses[0].ID
	// order.BillingAddressID = user.Addresses[0].ID

	// Order Item
	orderItem := models.OrderItem{}
	orderItem.OrderID = order.ID
	orderItem.Quantity = 1
	orderItem.Price = product.Price
	if err := DB(ctx).Create(&orderItem).Error; err != nil {
		fmt.Printf("create orderItem (%v) failure, got err %v", orderItem, err)
	}
	order.OrderItems = append(order.OrderItems, orderItem)

	order.PaymentAmount = order.Amount()
	// OrderStateMachine.Trigger("checkout", &order, DB(ctx), "test test test")

	if (user.Balance - order.PaymentAmount) >= 0 {
		order.State = "paid"
		user.Balance = user.Balance - order.PaymentAmount
		DB(ctx).Save(&user)
	} else {
		order.State = "cancelled"
	}

	if err := DB(ctx).Save(&order).Error; err != nil {
		fmt.Printf("Save order (%v) failure, got err %v", order, err)
	}

	fmt.Printf("name %v\n", product.Name)
	fmt.Printf("User %v\n", user.Balance)
	fmt.Printf("order %v\n", order.PaymentAmount)

}

func funcsMap(ctx *gin.Context) template.FuncMap {
	funcMaps := map[string]interface{}{
		"related_products": func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			DB(ctx).Preload("ColorVariations").Limit(4).Find(&products, "id <> ?", cv.ProductID)
			return products
		},
		"other_also_bought": func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			DB(ctx).Preload("ColorVariations").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
			return products
		},
	}
	for key, value := range I18nFuncMap(ctx) {
		funcMaps[key] = value
	}
	for key, value := range admin.ActionBar.FuncMap(ctx.Writer, ctx.Request) {
		funcMaps[key] = value
	}
	return funcMaps
}
