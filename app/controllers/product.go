package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
)

func ProductIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(10).Find(&products)
	seoObj := models.Seo{}
	db.DB.First(&seoObj)
	ctx.HTML(
		http.StatusOK,
		"product_index.tmpl",
		gin.H{
			"Products": products,
			"SeoTag":   seoObj.DefaultPage.Render(seoObj, nil),
			"MicroSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q=",
			}.Render(),
			"MicroContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
		},
	)
}

func ProductShow(ctx *gin.Context) {
	var product models.Product
	var colorVariation models.ColorVariation
	codes := strings.Split(ctx.Param("code"), "-")
	db.DB.Where(&models.Product{Code: codes[0]}).First(&product)
	db.DB.Preload("Images").Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: codes[1]}).First(&colorVariation)
	seoObj := models.Seo{}
	db.DB.First(&seoObj)

	var imageURL string
	if len(colorVariation.Images) > 0 {
		imageURL = colorVariation.Images[0].Image.URL()
	}

	ctx.HTML(
		http.StatusOK,
		"product_show.tmpl",
		gin.H{
			"Product":        product,
			"ColorVariation": colorVariation,
			"SeoTag":         seoObj.ProductPage.Render(seoObj, product),
			"MicroProduct": seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
				Image:       imageURL,
			}.Render(),
		},
	)
}
