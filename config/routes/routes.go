package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/controllers"
	"github.com/qor/qor-example/config"
	"html/template"
)

func Rounter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	if tmpl, err := template.New("projectViews").Funcs(config.FuncMap).ParseGlob("app/views/*.tmpl"); err == nil {
		router.SetHTMLTemplate(tmpl)
	} else {
		panic(err)
	}
	router.GET("/products", controllers.ProductIndex)
	router.GET("/products/:id", controllers.ProductShow)
	return router
}
