package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/config/seo"
	qor_seo "github.com/qor/seo"
	"github.com/qor/widget"
	"gopkg.in/authboss.v0"
)

func HomeIndex(ctx *gin.Context) {
	var products []models.Product
	DB(ctx).Limit(9).Preload("ColorVariations").Find(&products)

	widgetContext := admin.Widgets.NewContext(&widget.Context{
		DB:         DB(ctx),
		Options:    map[string]interface{}{"Request": ctx.Request},
		InlineEdit: IsEditMode(ctx),
	})

	config.View.Funcs(I18nFuncMap(ctx)).Execute(
		"home_index",
		gin.H{
			"ActionBarTag":           admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SeoCollection.SeoSettingURL("/help")}).Render(ctx.Writer, ctx.Request),
			authboss.FlashSuccessKey: auth.Auth.FlashSuccess(ctx.Writer, ctx.Request),
			authboss.FlashErrorKey:   auth.Auth.FlashError(ctx.Writer, ctx.Request),
			"SeoTag":                 seo.SeoCollection.Render(&qor.Context{DB: DB(ctx)}, "Default Page"),
			"top_banner":             widgetContext.Render("TopBanner", "Banner"),
			"feature_products":       widgetContext.Render("FeatureProducts", "Products"),
			"Products":               products,
			"MicroSearch": qor_seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q={keyword}",
			}.Render(),
			"MicroContact": qor_seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
