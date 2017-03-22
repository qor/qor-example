package cart

import (
	"fmt"
	"net/http"

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

	if err := ctx.BindJSON(&cartItem); err == nil {
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
	fmt.Printf("this is remove hendler %v\n", ctx)

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func ShowCartHandler(ctx *gin.Context) {
	// fmt.Printf("this is show %v\n", ctx)

	var (
		curCart, err   = GetCart(ctx)
		cartItems      = curCart.GetContent()
		sizeVariations []models.SizeVariation
		itemIDS        = make([]uint, 0, len(cartItems))
		extCartItems   []fullCartItem
	)

	curCart.Each(func(item *CartItem, key uint) {
		itemIDS = append(itemIDS, key)
	})

	DB(ctx).Preload("ColorVariation.Color").Preload("ColorVariation.Product").Where(itemIDS).Find(&sizeVariations)

	for _, item := range sizeVariations {
		extCartItems = append(extCartItems, fullCartItem{
			*cartItems[item.ID],
			item,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":    "OK",
		"cartItems": extCartItems,
		"message":   "Found cart items",
		"count":     len(extCartItems),
	})

}
