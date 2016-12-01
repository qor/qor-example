package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/db"
	"github.com/qor/qor/utils"
	"github.com/qor/wildcard_router"
)

var rootMux *http.ServeMux
var WildcardRouter *wildcard_router.WildcardRouter

func Router() *http.ServeMux {
	if rootMux == nil {
		router := gin.Default()
		router.Use(func(ctx *gin.Context) {
			tx := db.DB
			if locale := utils.GetLocale(&qor.Context{Request: ctx.Request, Writer: ctx.Writer}); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			if publishScheduledTime := publish2.GetScheduledTime(ctx.Request, ctx.Writer); publishScheduledTime != "" {
				if t, err := utils.ParseTime(publishScheduledTime, &qor.Context{Request: ctx.Request, Writer: ctx.Writer}); err == nil {
					tx = tx.Set(publish2.ScheduledTime, t)
				}
			}

			ctx.Set("DB", tx)
		})

		gin.SetMode(gin.DebugMode)

		router.GET("/", controllers.HomeIndex)
		router.GET("/products/:code", controllers.ProductShow)
		router.GET("/switch_locale", controllers.SwitchLocale)

		rootMux = http.NewServeMux()
		rootMux.Handle("/auth/", auth.Auth.NewRouter())
		publicDir := http.Dir(strings.Join([]string{config.Root, "public"}, "/"))
		rootMux.Handle("/dist/", utils.FileServer(publicDir))
		rootMux.Handle("/vendors/", utils.FileServer(publicDir))
		rootMux.Handle("/images/", utils.FileServer(publicDir))
		rootMux.Handle("/fonts/", utils.FileServer(publicDir))

		WildcardRouter = wildcard_router.New()
		WildcardRouter.MountTo("/", rootMux)
		WildcardRouter.AddHandler(router)
	}
	return rootMux
}
