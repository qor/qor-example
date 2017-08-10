package controllers

import (
    "net/http"
    // "strings"

    // "github.com/qor/qor-example/app/models"
    "github.com/qor/qor-example/config"
    // "github.com/qor/qor-example/config/utils"
)

func ShowCartHandler(w http.ResponseWriter, req *http.Request) {
    // var (
    //     product        models.Product
        // colorVariation models.ColorVariation
        // codes          = strings.Split(utils.URLParam("code", req), "_")
    //     productCode    = codes[0]
        // colorCode      string
    //     tx             = utils.GetDB(req)
    // )

    // if len(codes) > 1 {
    //     colorCode = codes[1]
    // }

    // if tx.Preload("Category").Where(&models.Product{Code: productCode}).First(&product).RecordNotFound() {
    //     http.Redirect(w, req, "/", http.StatusFound)
    // }

    // tx.Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)

    config.View.Execute("/cart/cart_show", map[string]interface{}{}, req, w)
}


func CheckoutCartHandler(w http.ResponseWriter, req *http.Request) {
    // var (
    //     product        models.Product
        // colorVariation models.ColorVariation
    //     codes          = strings.Split(utils.URLParam("code", req), "_")
    //     productCode    = codes[0]
    //     colorCode      string
    //     tx             = utils.GetDB(req)
    // )

    // if len(codes) > 1 {
    //     colorCode = codes[1]
    // }

    // if tx.Preload("Category").Where(&models.Product{Code: productCode}).First(&product).RecordNotFound() {
    //     http.Redirect(w, req, "/", http.StatusFound)
    // }

    // tx.Preload("Product").Preload("Color").Preload("SizeVariations.Size").Where(&models.ColorVariation{ProductID: product.ID, ColorCode: colorCode}).First(&colorVariation)

    config.View.Execute("/cart/cart_checkout", map[string]interface{}{}, req, w)
}