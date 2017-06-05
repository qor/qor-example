package routes

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"

	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/db"
)

var rootMux *http.ServeMux
var WildcardRouter *wildcard_router.WildcardRouter

func Router() *http.ServeMux {
	if rootMux == nil {
		router := gin.Default()

		store := sessions.NewCookieStore([]byte("something-very-secret"))
		router.Use(sessions.Sessions("mysession", store))

		router.Use(func(ctx *gin.Context) {
			tx := db.DB
			context := &qor.Context{Request: ctx.Request, Writer: ctx.Writer}
			if locale := utils.GetLocale(context); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			ctx.Set("DB", publish2.PreviewByDB(tx, context))
		})

		gin.SetMode(gin.DebugMode)

		router.GET("/", controllers.HomeIndex)
		router.GET("/products/:code", controllers.ProductShow)
		router.GET("/category/:code", controllers.CategoryShow)
		router.GET("/switch_locale", controllers.SwitchLocale)

		cabinetGroup := router.Group("/cabinet")
		{
			cabinetGroup.GET("/", controllers.CabinetShow)
			cabinetGroup.POST("/add_user_credit", controllers.AddUserCredit)
			cabinetGroup.GET("/profile", controllers.ProfileShow)
			cabinetGroup.POST("/profile", controllers.SetUserProfile)
			cabinetGroup.POST("/profile/billing_address", controllers.SetBillingAddress)
			cabinetGroup.POST("/profile/shipping_address", controllers.SetShippingAddress)
		}

		cartGroup := router.Group("/cart")
		{
			cartGroup.GET("/", controllers.ShowCartHandler)
			cartGroup.GET("/checkout", controllers.CheckoutCartHandler)
			cartGroup.POST("/", controllers.AddToCartHandler)
			cartGroup.POST("/checkout", controllers.OrderCartHandler)
			cartGroup.DELETE("/:id", controllers.RemoveFromCartHandler)
		}

		rootMux = http.NewServeMux()

		rootMux.Handle("/auth/", auth.Auth.NewServeMux("/auth"))

		rootMux.Handle("/system/", utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))))
		assetFS := bindatafs.AssetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors")
		for _, path := range []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors"} {
			rootMux.Handle(fmt.Sprintf("/%s/", path), assetFS)
		}

		WildcardRouter = wildcard_router.New()
		WildcardRouter.MountTo("/", rootMux)
		WildcardRouter.AddHandler(router)
	}
	return rootMux
}
