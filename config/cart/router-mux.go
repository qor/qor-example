package cart

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	fmt.Printf("this is show %v\n", ctx)

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})

}
