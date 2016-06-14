package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor-example/config/i18n"
	"html/template"
	"net/http"
)

func SwitchLocale(ctx *gin.Context) {
	setCookie(http.Cookie{Name: "locale", Value: ctx.Request.URL.Query().Get("locale")}, ctx)
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

func I18nFuncMap(ctx *gin.Context) template.FuncMap {
	locale := "en-US"
	if cookie, err := ctx.Request.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return inline_edit.FuncMap(i18n.I18n, locale, isEditMode(ctx))
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
