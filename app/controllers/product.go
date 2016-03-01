package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
	"github.com/qor/qor-example/config"
	"github.com/apertoire/mlog"
)

// GET: http://localhost:7000/api/v1/category
func ProductApiIndex(ctx *gin.Context) {
	mlog.Start(mlog.LevelTrace, "")
	var products []models.Product
	acceptLanguage := ctx.Request.Header.Get("Accept-Language")[0:2]
	locale := ctx.Request.Header.Get("Locale")
	if len(locale) == 0 {
		locale = config.Config.Locale
	}
	mlog.Trace("acceptLanguage: %v, locale: %v", acceptLanguage, locale)
	// session := sessions.Default(ctx)

	if err := db.DB.Set("l10n:locale", locale).Find(&products).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &products})
}

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
				Image:       colorVariation.MainImageUrl(),
			}.Render(),
		},
	)
}
