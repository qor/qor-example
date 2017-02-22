package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
)

func CategoryShow(ctx *gin.Context) {
	var (
		category models.Category
		products []models.Product
	)

	if DB(ctx).Where("code = ?", ctx.Param("code")).First(&category).RecordNotFound() {
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
