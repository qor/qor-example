package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/action_bar"
	"github.com/qor/qor"
	apputils "github.com/qor/qor-example/config/utils"
	"github.com/qor/qor/utils"
	"github.com/qor/widget"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/seo"
)

func HomeIndex(w http.ResponseWriter, req *http.Request) {
	var (
		products   []models.Product
		categories []models.Category
		tx         = apputils.GetDB(req)
	)

	tx.Limit(9).Preload("ColorVariations").Find(&products)
	tx.Find(&categories)

	widgetContext := admin.Widgets.NewContext(&widget.Context{
		DB:         tx,
		Options:    map[string]interface{}{"Request": req},
		InlineEdit: apputils.GetEditMode(w, req),
	})

	config.View.Execute(
		"home_index",
		gin.H{
			"ActionBarTag":     admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(w, req),
			"SEOTag":           seo.SEOCollection.Render(&qor.Context{DB: tx}, "Default Page"),
			"top_banner":       widgetContext.Render("TopBanner", "Banner"),
			"feature_products": widgetContext.Render("FeatureProducts", "Products"),
			"Products":         products,
		},
		req,
		w,
	)
}

func SwitchLocale(w http.ResponseWriter, req *http.Request) {
	utils.SetCookie(http.Cookie{Name: "locale", Value: req.URL.Query().Get("locale")}, &qor.Context{Request: req, Writer: w})
	http.Redirect(w, req, req.Referer(), http.StatusSeeOther)
}
