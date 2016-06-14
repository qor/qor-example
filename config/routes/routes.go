package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/auth"
)

func Router() *http.ServeMux {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/", controllers.HomeIndex)
	router.GET("/products/:code", controllers.ProductShow)
	router.GET("/switch_locale", controllers.SwitchLocale)

	var mux = http.NewServeMux()
	mux.Handle("/", router)
	mux.Handle("/auth/", auth.Auth.NewRouter())
	publicDir := http.Dir(strings.Join([]string{config.Root, "public"}, "/"))
	mux.Handle("/dist/", http.FileServer(publicDir))
	mux.Handle("/vendors/", http.FileServer(publicDir))
	mux.Handle("/images/", http.FileServer(publicDir))
	mux.Handle("/fonts/", http.FileServer(publicDir))
	return mux
}
