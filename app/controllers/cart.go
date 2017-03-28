package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qor/action_bar"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/qor-example/config/seo"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
)

func AddToCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = cart.GetCart(ctx)
		cartItem   cart.CartItem
	)

	if err := ctx.Bind(&cartItem); err == nil {
		curCart.Add(&cartItem)
		ctx.JSON(
			http.StatusCreated,
			gin.H{
				"status":  http.StatusCreated,
				"message": "Cart item added successfully!",
				"itemID":  cartItem.SizeVariationID,
			},
		)
	}
}

func RemoveFromCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = cart.GetCart(ctx)
		id, _      = strconv.Atoi(ctx.Param("id"))
	)

	fmt.Printf("%T, %v\n", id, id)
	curCart.Remove(uint(id))

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ShowCartHandler(ctx *gin.Context) {
	var (
		curCart, _     = cart.GetCart(ctx)
		cartItems      = curCart.GetContent()
		sizeVariations []models.SizeVariation
		itemIDS        = make([]uint, 0, len(cartItems))
		extCartItems   []cart.FullCartItem
	)

	curCart.Each(func(item *cart.CartItem, key uint) {
		itemIDS = append(itemIDS, key)
	})

	DB(ctx).Preload("Size").Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	var cartAmount, cartItemAmount float32
	for _, item := range sizeVariations {
		cartItemAmount = float32(item.ColorVariation.Product.Price) * float32(cartItems[item.ID].Quantity)
		cartAmount = cartAmount + cartItemAmount
		extCartItems = append(extCartItems, cart.FullCartItem{
			*cartItems[item.ID],
			item.ColorVariation.Product.MainImageURL(),
			item.ColorVariation.Product.Name,
			item.ColorVariation.Color.Name,
			item.Size.Name,
			item.ColorVariation.Product.Price,
			cartItemAmount,
		})
	}

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cart_show",
		gin.H{
			"ActionBarTag":  admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"showCartItems": extCartItems,
			"cartAmount":    cartAmount,
			"Categories":    CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)

	/* 	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Found %v cart items", len(extCartItems)),
		"data": gin.H{
			"items":  extCartItems,
			"count":  len(extCartItems),
			"amount": cartAmount,
		},
	}) */

}

func CheckoutCartHandler(ctx *gin.Context) {
	var (
		// curCart, _  = cart.GetCart(ctx)
		currentUser = CurrentUser(ctx)
		newOrder    models.Order
		addresses   []models.Address
	)

	// DB(ctx).Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	// curCart.Each(func(item *cart.CartItem) { item.Println(i.SizeVariationID) })

	if currentUser == nil {
		http.Redirect(ctx.Writer, ctx.Request, "/auth/login", http.StatusTemporaryRedirect)
	}

	DB(ctx).Model(&currentUser).Related(&addresses)

	fmt.Println(addresses)

	/* 	for addr := range currentUser.Addresses {
		fmt.Println(addr.Stringify())
	} */

	DB(ctx).Model(&newOrder)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"order_show",
		gin.H{
			"ActionBarTag": admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"Addresses":    addresses,
			// "cartAmount":    cartAmount,
			"Categories": CategoriesList(ctx),

			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
