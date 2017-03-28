package cart

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qor/action_bar"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/seo"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

func DB(ctx *gin.Context) *gorm.DB {
	newDB, exist := ctx.Get("DB")
	if exist {
		return newDB.(*gorm.DB)
	}
	return db.DB
}

func AddToCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = GetCart(ctx)
		cartItem   CartItem
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
		curCart, _ = GetCart(ctx)
		id, _      = strconv.Atoi(ctx.Param("id"))
	)

	fmt.Printf("%T, %v\n", id, id)
	curCart.Remove(uint(id))

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ShowCartHandler(ctx *gin.Context) {
	var (
		curCart, _     = GetCart(ctx)
		cartItems      = curCart.GetContent()
		sizeVariations []models.SizeVariation
		itemIDS        = make([]uint, 0, len(cartItems))
		extCartItems   []fullCartItem
	)

	curCart.Each(func(item *CartItem, key uint) {
		itemIDS = append(itemIDS, key)
	})

	DB(ctx).Preload("Size").Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	var cartAmount, cartItemAmount float32
	for _, item := range sizeVariations {
		cartItemAmount = float32(item.ColorVariation.Product.Price) * float32(cartItems[item.ID].Quantity)
		cartAmount = cartAmount + cartItemAmount
		extCartItems = append(extCartItems, fullCartItem{
			*cartItems[item.ID],
			item.ColorVariation.Product.MainImageURL(),
			item.ColorVariation.Product.Name,
			item.ColorVariation.Color.Name,
			item.Size.Name,
			item.ColorVariation.Product.Price,
			cartItemAmount,
		})
	}

	config.View.Funcs(controllers.I18nFuncMap(ctx)).Execute(
		"cart_show",
		gin.H{
			"ActionBarTag":  admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"showCartItems": extCartItems,
			"cartAmount":    cartAmount,
			"Categories":    controllers.CategoriesList(ctx),

			"CurrentUser":   controllers.CurrentUser(ctx),
			"CurrentLocale": controllers.CurrentLocale(ctx),
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
