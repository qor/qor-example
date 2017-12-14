package main

import (
	"github.com/qor/middlewares"
	"github.com/qor/qor-example/config/admin/bindatafs"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/session"

	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/qor/action_bar"
	"github.com/qor/i18n/inline_edit"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/admin"
	"github.com/qor/qor-example/config/api"
	"github.com/qor/qor-example/config/i18n"
	"github.com/qor/qor-example/config/routes"
	"github.com/qor/qor-example/config/seo"
	"github.com/qor/qor-example/config/utils"
	_ "github.com/qor/qor-example/db/migrations"
	"github.com/qor/render"
	"github.com/qor/session/manager"
	"github.com/qor/widget"
)

func main() {
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	cmdLine.Parse(os.Args[1:])

	mux := http.NewServeMux()
	mux.Handle("/", routes.Router())
	admin.Admin.MountTo("/admin", mux)
	admin.Filebox.MountTo("/downloads", mux)
	api.API.MountTo("/api", mux)

	config.View.FuncMapMaker = func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}

		// Add `t` method
		for key, fc := range inline_edit.FuncMap(i18n.I18n, utils.GetCurrentLocale(req), utils.GetEditMode(w, req)) {
			funcMap[key] = fc
		}

		for key, value := range admin.ActionBar.FuncMap(w, req) {
			funcMap[key] = value
		}

		widgetContext := admin.Widgets.NewContext(&widget.Context{
			DB:         utils.GetDB(req),
			Options:    map[string]interface{}{"Request": req},
			InlineEdit: utils.GetEditMode(w, req),
		})
		for key, fc := range widgetContext.FuncMap() {
			funcMap[key] = fc
		}

		funcMap["flashes"] = func() []session.Message {
			return manager.SessionManager.Flashes(w, req)
		}

		// Add `action_bar` method
		funcMap["render_action_bar"] = func() template.HTML {
			return admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(w, req)
		}

		funcMap["render_seo_tag"] = func() template.HTML {
			return seo.SEOCollection.Render(&qor.Context{DB: utils.GetDB(req)}, "Default Page")
		}

		funcMap["get_categories"] = func() (categories []models.Category) {
			utils.GetDB(req).Find(&categories)
			return
		}

		funcMap["current_locale"] = func() string {
			return utils.GetCurrentLocale(req)
		}

		funcMap["current_user"] = func() *models.User {
			return utils.GetCurrentUser(req)
		}

		funcMap["related_products"] = func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			utils.GetDB(req).Preload("ColorVariations").Limit(4).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["other_also_bought"] = func(cv models.ColorVariation) []models.Product {
			var products []models.Product
			utils.GetDB(req).Preload("ColorVariations").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["cart_qty"] = func() uint {
			curCart, _ := cart.GetCart(w, req)

			return uint(len(curCart.GetItemsIDS()))
		}

		funcMap["cart_list"] = func() (extCartItems []interface{}) {
			var (
				curCart, _ = cart.GetCart(w, req)
				svs        = models.SizeVariations()
			)

			utils.GetDB(req).Where(curCart.GetItemsIDS()).Find(&svs)

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
				svs        = models.SizeVariations()
			)

			utils.GetDB(req).Where(curCart.GetItemsIDS()).Find(&svs)

			amount = 0
			for _, sv := range svs {
				amount += float32(uint(sv.ColorVariation.Product.Price*100)*curCart.GetContent()[sv.ID].Quantity) / 100
			}

			return
		}

		return funcMap
	}

	if *compileTemplate {
		bindatafs.AssetFS.Compile()
	} else {
		fmt.Printf("Listening on: %v\n", config.Config.Port)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), middlewares.Apply(mux)); err != nil {
			panic(err)
		}
	}
}
