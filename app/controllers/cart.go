package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/qor-example/config/seo"
)

type showCartItem struct {
	SizeVariationID uint
	SizeVariation   models.SizeVariation
	Quantity        uint
	Amount          float32
	MainImageURL    string
}

func CartShow(ctx *gin.Context) {
	var (
		sizeVariations []models.SizeVariation
		showCartItems  []*showCartItem
		curCart, _     = cart.GetCart(ctx)
		cartItems      = curCart.GetContent()
	)

	itemIDS := make([]uint, 0, len(cartItems))
	for key, _ := range cartItems {
		itemIDS = append(itemIDS, key)
	}

	DB(ctx).Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	curCart.Each(func(i *cart.CartItem) { fmt.Println(i.SizeVariationID) })

	var cartAmount, cartItemAmount float32
	if !curCart.IsEmpty() {
		for _, variation := range sizeVariations {
			cartItemAmount = float32(variation.ColorVariation.Product.Price) * float32(cartItems[variation.ID].Quantity)
			cartAmount = cartAmount + cartItemAmount

			showCartItems = append(showCartItems, &showCartItem{
				SizeVariationID: cartItems[variation.ID].SizeVariationID,
				SizeVariation:   variation,
				Quantity:        cartItems[variation.ID].Quantity,
				Amount:          cartItemAmount,
				MainImageURL:    variation.ColorVariation.Product.MainImageURL(),
			})
			fmt.Printf("current cart item %v %v\n", variation.ID, cartItems[variation.ID])
		}
	}

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cart_show",
		gin.H{
			"ActionBarTag":  admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"showCartItems": showCartItems,
			"cartAmount":    cartAmount,
			"Categories":    CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
