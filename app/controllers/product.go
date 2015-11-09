package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"net/http"
)

func ProductIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(10).Find(&products)
	ctx.HTML(
		http.StatusOK,
		"product_index.tmpl",
		gin.H{
			"products": products,
		},
	)
}

func ProductShow(ctx *gin.Context) {
	var product models.Product
	db.DB.Find(&product, ctx.Param("id"))
	ctx.HTML(
		http.StatusOK,
		"product_show.tmpl",
		gin.H{
			"product": product,
		},
	)
}
