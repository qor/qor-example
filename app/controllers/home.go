package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
	"github.com/qor/widget"
)

func HomeIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Find(&products)
	seoObj := models.SEOSetting{}
	db.DB.First(&seoObj)

	i18nFuncMap := inline_edit.GenerateFuncMaps(i18n.I18n, "en-US", nil)
	widgetContext := widget.NewContext(map[string]interface{}{})
	config.View.Funcs(i18nFuncMap).Execute(
		"home_index",
		gin.H{
			"SeoTag":           seoObj.HomePage.Render(seoObj, nil),
			"top_banner":       admin.Widgets.Render("Banner", "TopBanner", widgetContext),
			"feature_products": admin.Widgets.Render("Products", "FeatureProducts", widgetContext),
			"Products":         products,
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
		ctx.Request,
		ctx.Writer,
	)
}
