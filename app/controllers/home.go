package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"

	"github.com/qor/widget"
)

func HomeIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Find(&products)
	seoObj := models.SEOSetting{}
	db.DB.First(&seoObj)

	widgetContext := widget.NewContext(map[string]interface{}{})
	ctx.HTML(
		http.StatusOK,
		"home_index.tmpl",
		gin.H{
			"SeoTag":         seoObj.HomePage.Render(seoObj, nil),
			"banner_widget":  admin.Widget.Render("HomeBanner", widgetContext, "Banner"),
			"banner_widget1": admin.Widget.Render("HomeBanner1", widgetContext, "Banner"),
			"Products":       products,
			"MicroSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q={keyword}",
			}.Render(),
			"MicroContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
		},
	)
}
