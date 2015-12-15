package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
	"net/http"
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
			"products": products,
			"seoTag":   seoObj.DefaultPage.Render(seoObj, nil),
			"microSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q=",
			}.Render(),
			"microContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
		},
	)
}

func ProductShow(ctx *gin.Context) {
	var product models.Product
	db.DB.Find(&product, ctx.Param("id"))
	seoObj := models.Seo{}
	db.DB.First(&seoObj)
	ctx.HTML(
		http.StatusOK,
		"product_show.tmpl",
		gin.H{
			"product": product,
			"seoTag":  seoObj.ProductPage.Render(seoObj, product),
			"microProduct": seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
			}.Render(),
		},
	)
}
