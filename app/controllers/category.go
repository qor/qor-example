package controllers

import (
	"net/http"

	"github.com/apertoire/mlog"
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	// "github.com/gin-gonic/contrib/sessions"
)

// GET: http://localhost:7000/api/v1/category
func CategoryIndex(ctx *gin.Context) {
	mlog.Start(mlog.LevelTrace, "")
	var categorys []models.Category
	acceptLanguage := ctx.Request.Header.Get("Accept-Language")[0:2]
	locale := ctx.Request.Header.Get("Locale")
	if len(locale) == 0 {
		locale = config.Config.Locale
	}
	mlog.Trace("acceptLanguage: %v, locale: %v", acceptLanguage, locale)
	// session := sessions.Default(ctx)

	if err := db.DB.Set("l10n:locale", locale).Find(&categorys).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &categorys})
}
