package controllers

import (
	"fmt"

	// "html/template"
	"net/http"
	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/qor/action_bar"
	// "github.com/qor/qor"
	// "github.com/jinzhu/gorm"
	// "github.com/qor/slug"
	"github.com/qor/qor-example/app/models"
	// "github.com/qor/qor-example/config"
	// "github.com/qor/qor-example/config/admin"
	// "github.com/qor/qor-example/config/seo"
	// qor_seo "github.com/qor/seo"
)

func CategoryShow(ctx *gin.Context) {
	var (
		category models.Category
	)
	if DB(ctx).Where("name_with_slug = ?", ctx.Param("name")).First(&category).RecordNotFound() {
		http.Redirect(ctx.Writer, ctx.Request, "/", http.StatusFound)
	}

	fmt.Printf(ctx.Param("name") + "\n" + category.NameWithSlug.Slug + "\n ololo\n")
}
