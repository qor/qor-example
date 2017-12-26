package orders

import (
	"net/http"

	"github.com/qor/render"
)

// Controller products controller
type Controller struct {
	View *render.Render
}

// Cart shopping cart
func (ctrl Controller) Cart(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("cart", map[string]interface{}{}, req, w)
}

// UpdateCart update shopping cart
func (ctrl Controller) UpdateCart(w http.ResponseWriter, req *http.Request) {
	// FIXME
}

// Checkout checkout shopping cart
func (ctrl Controller) Checkout(w http.ResponseWriter, req *http.Request) {
	// FIXME
}
