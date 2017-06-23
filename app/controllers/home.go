package controllers

import (
	"net/http"

	"github.com/qor/qor"
	apputils "github.com/qor/qor-example/config/utils"
	"github.com/qor/qor/utils"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	var (
		tx       = apputils.GetDB(req)
		products []models.Product
	)

	tx.Limit(9).Preload("ColorVariations").Find(&products)

	config.View.Execute(
		"home_index",
		map[string]interface{}{
			"Products": products,
		},
		req,
		w,
	)
}

func SwitchLocale(w http.ResponseWriter, req *http.Request) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: req.URL.Query().Get("locale")}, &qor.Context{Request: req, Writer: w})
	http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
}
