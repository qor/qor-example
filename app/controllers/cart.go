package controllers

import (
	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/qor-example/config/seo"
)

func CartShow(ctx *gin.Context) {
	var (
		// order          = CurrentOrder(ctx)
		// orderItems     []models.OrderItem
		sizeVariations []models.SizeVariation
		curCart        *cart.Cart
	)

	curCart, _ = cart.GetCart(cart.GinGonicSession{sessions.Default(ctx)})
	itemIDS := make([]uint, 0, len(curCart.GetContent()))
	for key, item := range curCart.GetContent() {
		itemIDS = append(itemIDS, key)
		fmt.Printf("cart item %v %v\n", key, item)
	}

	fmt.Println(itemIDS)

	DB(ctx).Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	// DB(ctx).Model(&order).Preload("SizeVariation.ColorVariation.Color").Preload("SizeVariation.ColorVariation.Product").Where(&models.OrderItem{OrderID: order.ID}).Find(&orderItems)
	// fmt.Printf("CartShow Order: %v\n", order.ID)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cart_show",
		gin.H{
			"ActionBarTag": admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			// "OrderItems":     orderItems,
			"SizeVariations": sizeVariations,
			// "Order":          order,
			"Categories": CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
