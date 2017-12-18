package controllers

import (
	"net/http"

	"github.com/qor/qor"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor/utils"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	config.View.Execute("home_index", map[string]interface{}{}, req, w)
}

func SwitchLocale(w http.ResponseWriter, req *http.Request) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: req.URL.Query().Get("locale")}, &qor.Context{Request: req, Writer: w})
	http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
}
