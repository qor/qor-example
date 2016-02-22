package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

func HomeIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Find(&products)

	ctx.HTML(
		http.StatusOK,
		"home_index.tmpl",
		gin.H{
			"products": products,
		},
	)
}
