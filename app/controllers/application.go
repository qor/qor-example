package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/auth"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/db"
)

func SwitchLocale(ctx *gin.Context) {
	setCookie(http.Cookie{Name: "locale", Value: ctx.Request.URL.Query().Get("locale")}, ctx)
	ctx.Redirect(http.StatusTemporaryRedirect, ctx.Request.Referer())
}

func CurrentLocale(ctx *gin.Context) string {
	locale := "en-US"
	if cookie, err := ctx.Request.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

func CurrentUser(ctx *gin.Context) *models.User {
	userInter, err := auth.Auth.CurrentUser(ctx.Writer, ctx.Request)
	if userInter != nil && err == nil {
		return userInter.(*models.User)
	}
	return nil
}

func IsEditMode(ctx *gin.Context) bool {
	return admin.ActionBar.EditMode(ctx.Writer, ctx.Request)
}

func I18nFuncMap(ctx *gin.Context) template.FuncMap {
	return inline_edit.FuncMap(i18n.I18n, CurrentLocale(ctx), IsEditMode(ctx))
}

func setCookie(cookie http.Cookie, context *gin.Context) {
	cookie.HttpOnly = true

	// set https cookie
	if context.Request != nil && context.Request.URL.Scheme == "https" {
		cookie.Secure = true
	}

	// set default path
	if cookie.Path == "" {
		cookie.Path = "/"
	}

	http.SetCookie(context.Writer, &cookie)
}

func DB(ctx *gin.Context) *gorm.DB {
	newDB, exist := ctx.Get("DB")
	if exist {
		return newDB.(*gorm.DB)
	}
	return db.DB
}
