package controllers

import (
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/db"
	"github.com/qor/seo"
	"github.com/qor/widget"
	"gopkg.in/authboss.v0"
)

func IsEnableInlineEdit(ctx *gin.Context) bool {
	return admin.ActionBar.IsChecked(ctx.Writer, ctx.Request)
}

func CurrentUser(ctx *gin.Context) *models.User {
	userInter, err := auth.Auth.CurrentUser(ctx.Writer, ctx.Request)
	if userInter != nil && err == nil {
		return userInter.(*models.User)
	}
	return nil
}

func I18nFuncMap(ctx *gin.Context) template.FuncMap {
	return inline_edit.FuncMap(i18n.I18n, "en-US", IsEnableInlineEdit(ctx))
}

func HomeIndex(ctx *gin.Context) {
	var products []models.Product
	db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Find(&products)
	seoObj := models.SEOSetting{}
	db.DB.First(&seoObj)

	widgetContext := widget.NewContext(map[string]interface{}{"Request": ctx.Request})

	config.View.Funcs(I18nFuncMap(ctx)).Execute(
		"home_index",
		gin.H{
			"ActionBarTag":           admin.ActionBar.RenderIncludedTag(ctx.Writer, ctx.Request),
			authboss.FlashSuccessKey: auth.Auth.FlashSuccess(ctx.Writer, ctx.Request),
			authboss.FlashErrorKey:   auth.Auth.FlashError(ctx.Writer, ctx.Request),
			"SeoTag":                 seoObj.HomePage.Render(seoObj, nil),
			"top_banner":             admin.Widgets.Render("TopBanner", "Banner", widgetContext, IsEnableInlineEdit(ctx)),
			"feature_products":       admin.Widgets.Render("FeatureProducts", "Products", widgetContext, IsEnableInlineEdit(ctx)),
			"Products":               products,
			"MicroSearch": seo.MicroSearch{
				URL:    "http://demo.getqor.com",
				Target: "http://demo.getqor.com/search?q={keyword}",
			}.Render(),
			"MicroContact": seo.MicroContact{
				URL:         "http://demo.getqor.com",
				Telephone:   "080-0012-3232",
				ContactType: "Customer Service",
			}.Render(),
			"CurrentUser": CurrentUser(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
