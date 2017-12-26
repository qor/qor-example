package utils

import (
	"html/template"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"github.com/qor/action_bar"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/admin"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/qor-example/models/seo"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/render"
	"github.com/qor/session"
	"github.com/qor/session/manager"
	"github.com/qor/widget"
)

// HTMLSanitizer HTML sanitizer
var HTMLSanitizer = bluemonday.UGCPolicy()

// AddFuncMapMaker add FuncMapMaker to view
func AddFuncMapMaker(view *render.Render) *render.Render {
	oldFuncMapMaker := view.FuncMapMaker
	view.FuncMapMaker = func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			funcMap = oldFuncMapMaker(render, req, w)
		}

		// Add `t` method
		for key, fc := range inline_edit.FuncMap(i18n.I18n, GetCurrentLocale(req), GetEditMode(w, req)) {
			funcMap[key] = fc
		}

		for key, value := range admin.ActionBar.FuncMap(w, req) {
			funcMap[key] = value
		}

		widgetContext := admin.Widgets.NewContext(&widget.Context{
			DB:         GetDB(req),
			Options:    map[string]interface{}{"Request": req},
			InlineEdit: GetEditMode(w, req),
		})
		for key, fc := range widgetContext.FuncMap() {
			funcMap[key] = fc
		}

		funcMap["raw"] = func(str string) template.HTML {
			return template.HTML(HTMLSanitizer.Sanitize(str))
		}

		funcMap["flashes"] = func() []session.Message {
			return manager.SessionManager.Flashes(w, req)
		}

		// Add `action_bar` method
		funcMap["render_action_bar"] = func() template.HTML {
			return admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(w, req)
		}

		funcMap["render_seo_tag"] = func() template.HTML {
			return seo.SEOCollection.Render(&qor.Context{DB: GetDB(req)}, "Default Page")
		}

		funcMap["get_categories"] = func() (categories []products.Category) {
			GetDB(req).Find(&categories)
			return
		}

		funcMap["current_locale"] = func() string {
			return GetCurrentLocale(req)
		}

		funcMap["current_user"] = func() *users.User {
			return GetCurrentUser(req)
		}

		funcMap["related_products"] = func(cv products.ColorVariation) []products.Product {
			var products []products.Product
			GetDB(req).Preload("ColorVariations").Limit(4).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["other_also_bought"] = func(cv products.ColorVariation) []products.Product {
			var products []products.Product
			GetDB(req).Preload("ColorVariations").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["cart_qty"] = func() uint {
			curCart, _ := cart.GetCart(w, req)
			return uint(len(curCart.GetItemsIDS()))
		}

		funcMap["cart_list"] = func() (extCartItems []interface{}) {
			var (
				curCart, _ = cart.GetCart(w, req)
				svs        = products.SizeVariations()
			)

			GetDB(req).Where(curCart.GetItemsIDS()).Find(&svs)

			for _, sv := range svs {
				amount := float32(uint(sv.ColorVariation.Product.Price*100)*curCart.GetContent()[sv.ID].Quantity) / 100
				cartItem := curCart.GetContent()[sv.ID]

				extCartItems = append(extCartItems, map[string]interface{}{
					"GetImageURL": sv.ColorVariation.Product.MainImageURL(),
					"Name":        sv.ColorVariation.Product.Name,
					"Price":       sv.ColorVariation.Product.Price,
					"Amount":      amount,
					"DefaultPath": sv.ColorVariation.Product.DefaultPath(),

					"ProductID": cartItem.ProductID,
					"Quantity":  cartItem.Quantity,
				})
			}
			return
		}

		funcMap["cart_amount"] = func() (amount float32) {
			var (
				curCart, _ = cart.GetCart(w, req)
				svs        = products.SizeVariations()
			)

			GetDB(req).Where(curCart.GetItemsIDS()).Find(&svs)

			amount = 0
			for _, sv := range svs {
				amount += float32(uint(sv.ColorVariation.Product.Price*100)*curCart.GetContent()[sv.ID].Quantity) / 100
			}

			return
		}

		return funcMap
	}

	return view
}
