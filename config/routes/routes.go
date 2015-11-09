package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/controllers"
)

func Rounter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.LoadHTMLGlob("app/views/*.tmpl")
	router.GET("/products", controllers.ProductIndex)
	router.GET("/products/:id", controllers.ProductShow)
	return router
}
