package controllers

import (
	// "html/template"
	"net/http"
	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	// "github.com/qor/qor"
	// "github.com/jinzhu/gorm"
	// "github.com/qor/slug"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	// "github.com/qor/qor-example/config/seo"
	// qor_seo "github.com/qor/seo"
)

func CategoryShow(ctx *gin.Context) {
	var (
		category models.Category
		products []models.Product
	)
	if DB(ctx).Where("name_with_slug = ?", ctx.Param("name")).First(&category).RecordNotFound() {
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusFound)
	}
	DB(ctx).Where(&models.Product{CategoryID: category.ID}).Preload("ColorVariations").Find(&products)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"category_show",
		gin.H{
			"ActionBarTag":  admin.ActionBar.Actions(action_bar.EditResourceAction{Value: category, Inline: true, EditModeOnly: true}).Render(ctx.Writer, ctx.Request),
			"CurrentUser":   CurrentUser(ctx),
			"CurrentLocale": CurrentLocale(ctx),
			"CategoryName":  category.Name,
			"Products":      products,
			"Categories":    CategoriesList(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}
