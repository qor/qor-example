package cart

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/utils"
	// "github.com/qor/qor-example/app/controllers"
	// "github.com/qor/qor-example/app/models"
	// "github.com/qor/qor-example/config"
)

var RouterMux *gin.Engine

func init() {
	RouterMux = gin.Default()
	store := sessions.NewCookieStore([]byte("something-very-secret"))
	RouterMux.Use(func(ctx *gin.Context) {
		tx := db.DB
		context := &qor.Context{Request: ctx.Request, Writer: ctx.Writer}
		if locale := utils.GetLocale(context); locale != "" {
			tx = tx.Set("l10n:locale", locale)
		}

		ctx.Set("DB", publish2.PreviewByDB(tx, context))
	})
	RouterMux.Use(sessions.Sessions("mysession", store))

	RouterMux.GET("/cart", showCartHandler)
	RouterMux.POST("/cart/add", addToCart)
	RouterMux.GET("/cart/delete/:id", removeFromCart)
}

func Router() *gin.Engine {
	fmt.Println("returning route mux")
	return RouterMux

}

func addToCart(ctx *gin.Context) {
	var (
		curCart, _ = GetCart(ctx)
		cartItem   CartItem
	)

	ctx.Bind(&cartItem)
	curCart.Add(&cartItem)

	fmt.Println(cartItem)

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func removeFromCart(ctx *gin.Context) {
	fmt.Printf("this is remove hendler %v\n", ctx)

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func showCartHandler(ctx *gin.Context) {
	fmt.Printf("this is show %v\n", ctx)

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})

}
