package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor"
	"github.com/qor/widget"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/seo"
)

func HomeIndex(ctx *gin.Context) {
	var (
		products   []models.Product
		categories []models.Category
	)
	DB(ctx).Limit(9).Preload("ColorVariations").Find(&products)
	DB(ctx).Find(&categories)

	widgetContext := admin.Widgets.NewContext(&widget.Context{
		DB:         DB(ctx),
		Options:    map[string]interface{}{"Request": ctx.Request},
		InlineEdit: IsEditMode(ctx),
	})

	config.View.Funcs(I18nFuncMap(ctx)).Execute(
		"home_index",
		gin.H{
			"ActionBarTag":     admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			"SEOTag":           seo.SEOCollection.Render(&qor.Context{DB: DB(ctx)}, "Default Page"),
			"top_banner":       widgetContext.Render("TopBanner", "Banner"),
			"feature_products": widgetContext.Render("FeatureProducts", "Products"),
			"Products":         products,
			"Categories":       CategoriesList(ctx),
			"CurrentUser":      CurrentUser(ctx),
			"CurrentLocale":    CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
