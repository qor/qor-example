package controllers

import (
	"net/http"
	"strings"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/utils"
)

func ProductIndex(w http.ResponseWriter, req *http.Request) {
	var (
		products        []models.Product
		tx              = utils.GetDB(req)
	)

	tx.Preload("Category").Find(&products)

	config.View.Execute("/product/product", map[string]interface{}{"Products": products}, req, w)
}

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

	config.View.Execute("/product/product_details", map[string]interface{}{"CurrentColorVariation": colorVariation}, req, w)
}

func ProductGenderShow(w http.ResponseWriter, req *http.Request) {
	var (
		products []models.Product
		tx       = utils.GetDB(req)
	)

	tx.Where(&models.Product{Gender: utils.URLParam("gender", req)}).Preload("Category").Find(&products)

	config.View.Execute("/product/gender_list", map[string]interface{}{"Products": products}, req, w)
}
