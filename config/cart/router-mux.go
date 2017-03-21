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
		mux.GET("/cart/", showCartHandler)
		mux.POST("/cart/add", addToCartHandler)
		mux.GET("/cart/delete/:id", removeFromCart)
	}

	return mux

}

func addToCartHandler(ctx *gin.Context) {
	fmt.Printf("this is add hendler %v\n", ctx)

	ctx.JSON(
		http.StatusOK,
		"{success: true}",
	)
}

func removeFromCart(ctx *gin.Context) {
	fmt.Printf("this is remove hendler %v\n", ctx)

	ctx.JSON(
		http.StatusOK,
		"{success: true}",
	)
}

func showCartHandler(ctx *gin.Context) {
	fmt.Printf("this is show %v\n", ctx)

	ctx.JSON(
		http.StatusOK,
		"{somegood: true}",
	)
}
