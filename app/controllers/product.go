package controllers

import (
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
)

func ProductShow(ctx *gin.Context) {
	var (
		product        models.Product
		colorVariation models.ColorVariation
		seoSetting     models.SEOSetting
		codes          = strings.Split(ctx.Param("code"), "_")
		productCode    = codes[0]
		colorCode      string
	)

	if len(codes) > 1 {
		colorCode = codes[1]
	}

	db.DB.Where(&models.Product{Code: productCode}).First(&product)
	db.DB.Preload("Images").Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)
	db.DB.First(&seoSetting)

	config.View.Funcs(funcsMap()).Execute(
		"product_show",
		gin.H{
			"Product":        product,
			"ColorVariation": colorVariation,
			"SeoTag":         seoSetting.ProductPage.Render(seoSetting, product),
			"MicroProduct": seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
				Image:       colorVariation.MainImageUrl(),
			}.Render(),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func funcsMap() template.FuncMap {
	productFuncsMap := inline_edit.GenerateFuncMaps(i18n.I18n, "en-US", nil)
	productFuncsMap["related_products"] = func(cv models.ColorVariation) []models.Product {
		var products []models.Product
		db.DB.Preload("ColorVariations").Preload("ColorVariations.Images").Limit(4).Find(&products, "id <> ?", cv.ProductID)
		return products
	}
	productFuncsMap["other_also_bought"] = func(cv models.ColorVariation) []models.Product {
		var products []models.Product
		db.DB.Preload("ColorVariations").Preload("ColorVariations.Images").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
		return products
	}
	return productFuncsMap
}
