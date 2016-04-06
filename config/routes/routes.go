package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
)

func Router() *http.ServeMux {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/", controllers.HomeIndex)
	router.GET("/products/:code", controllers.ProductShow)

	var mux = http.NewServeMux()
	mux.Handle("/", router)
	publicDir := http.Dir(strings.Join([]string{config.Root, "public"}, "/"))
	mux.Handle("/public/", http.StripPrefix("/public", http.FileServer(publicDir)))
	mux.Handle("/images/", http.FileServer(publicDir))
	mux.Handle("/js/", http.FileServer(publicDir))
	mux.Handle("/css/", http.FileServer(publicDir))
	mux.Handle("/fonts/", http.FileServer(publicDir))
	return mux
}
