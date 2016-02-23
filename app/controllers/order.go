package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	// "github.com/gin-gonic/contrib/sessions"
)

func OrderIndex(ctx *gin.Context) {
	var orders []models.Order
	// session := sessions.Default(ctx)

	if err := db.DB.Limit(50).Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})

	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &orders})

	}

	// ctx.HTML(
	// 	http.StatusOK,
	// 	"product_index.tmpl",
	// 	gin.H{
	// 		"orders": orders,
	// 		"seoTag":   seoObj.DefaultPage.Render(seoObj, nil),
	// 		"microSearch": seo.MicroSearch{
	// 			URL:    "http://demo.getqor.com",
	// 			Target: "http://demo.getqor.com/search?q=",
	// 		}.Render(),
	// 		"microContact": seo.MicroContact{
	// 			URL:         "http://demo.getqor.com",
	// 			Telephone:   "080-0012-3232",
	// 			ContactType: "Customer Service",
	// 		}.Render(),
	// 	},
	// )
}

// func OrderShow(ctx *gin.Context) {
// 	var order models.Order
// 	db.DB.Preload("ColorVariations").Preload("ColorVariations.Images").Find(&order, ctx.Param("id"))
// 	seoObj := models.Seo{}
// 	db.DB.First(&seoObj)

// 	var imageURL string
// 	if len(order.ColorVariations) > 0 && len(order.ColorVariations[0].Images) > 0 {
// 		imageURL = order.ColorVariations[0].Images[0].Image.URL()
// 	}

// 	ctx.HTML(
// 		http.StatusOK,
// 		"order_show.tmpl",
// 		gin.H{
// 			"order": order,
// 			"seoTag":  seoObj.ProductPage.Render(seoObj, order),
// 			"microProduct": seo.MicroProduct{
// 				Name:        product.Name,
// 				Description: product.Description,
// 				BrandName:   product.Category.Name,
// 				SKU:         product.Code,
// 				Price:       float64(product.Price),
// 				Image:       imageURL,
// 			}.Render(),
// 		},
// 	)
// }
