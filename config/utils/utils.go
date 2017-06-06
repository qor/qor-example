package utils

import (
	"net/http"

	"github.com/qor/l10n"
	"github.com/qor/qor-example/config/admin"
)

// GetCurrentLocale get current locale from request
func GetCurrentLocale(request *http.Request) string {
	locale := l10n.Global
	if cookie, err := request.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

// GetEditMode get edit mode
func GetEditMode(writer http.ResponseWriter, req *http.Request) bool {
	return admin.ActionBar.EditMode(writer, req)
}
