package controllers

import (
	"net/http"
	"strings"

	qor_seo "github.com/qor/seo"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/utils"
)

func ProductShow(w http.ResponseWriter, req *http.Request) {
	var (
		product        models.Product
		colorVariation models.ColorVariation
		codes          = strings.Split(utils.URLParam("code", req), "_")
		productCode    = codes[0]
		colorCode      string
		tx             = utils.GetDB(req)
	)

	if len(codes) > 1 {
		colorCode = codes[1]
	}

	if tx.Preload("Category").Where(&models.Product{Code: productCode}).First(&product).RecordNotFound() {
		http.Redirect(w, req, "/", http.StatusFound)
	}

	tx.Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)

	config.View.Execute(
		"product_show",
		map[string]interface{}{
			"Product":        product,
			"ColorVariation": colorVariation,
			"MicroProduct": qor_seo.MicroProduct{
				Name:        product.Name,
				Description: product.Description,
				BrandName:   product.Category.Name,
				SKU:         product.Code,
				Price:       float64(product.Price),
				Image:       colorVariation.MainImageURL(),
			}.Render(),
		},
		req,
		w,
	)
}
