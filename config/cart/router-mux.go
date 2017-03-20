package cart

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	// "github.com/qor/qor"
	// "github.com/qor/qor-example/app/controllers"
	// "github.com/qor/qor-example/app/models"
	// "github.com/qor/qor-example/config"
	// "github.com/qor/qor-example/db"
)

var mux *gin.Engine

func NewRouter() *gin.Engine {
	fmt.Println("returning route mux")

	if mux == nil {
		mux = gin.New()
		mux.GET("/cart/", showHandler)
		mux.POST("/cart/add", showHandler)
	}

	return mux

}

func addToCartHandler(ctx *gin.Context) {
	var (
		curCart, _ = cart.GetCart(ctx)
		cartItem   cart.CartItem
	)
}

func showHandler(ctx *gin.Context) {
	fmt.Printf("this is cart mux %v\n", ctx)

	ctx.JSON(
		http.StatusOK,
		"{somegood: true}",
	)
}
