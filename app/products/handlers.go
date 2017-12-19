package products

import (
	"net/http"

	"github.com/qor/app/modules/products"
	"github.com/qor/qor-example/utils"
	"github.com/qor/render"
)

// Controller products controller
type Controller struct {
	View *render.Render
}

// Index products index page
func (ctrl Controller) Index(w http.ResponseWriter, req *http.Request) {
	var (
		products []products.Product
		tx       = utils.GetDB(req)
	)

	tx.Preload("Category").Find(&products)

	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
}

// Show product show page
func (ctrl Controller) Show(w http.ResponseWriter, req *http.Request) {
	var (
		products []products.Product
		tx       = utils.GetDB(req)
	)

	tx.Preload("Category").Find(&products)

	ctrl.View.Execute("index", map[string]interface{}{}, req, w)
}

// Gender products gender page
func (ctrl Controller) Gender(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("gender", map[string]interface{}{}, req, w)
}
