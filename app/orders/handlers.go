package orders

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/models/orders"
	"github.com/qor/qor-example/utils"
	qorrender "github.com/qor/render"
	"github.com/qor/responder"
	"github.com/qor/session/manager"
)

// Controller products controller
type Controller struct {
	View *qorrender.Render
}

var decoder = schema.NewDecoder()

// Cart shopping cart
func (ctrl Controller) Cart(w http.ResponseWriter, req *http.Request) {
	order := getCurrentOrder(w, req)
	tx := utils.GetDB(req)

	tx.Model(order).Association("OrderItems").Find(&order.OrderItems)

	ctrl.View.Execute("cart", map[string]interface{}{"Order": order}, req, w)
}

type updateCartInput struct {
	SizeVariationID uint `schema:"size_variation_id"`
	Quantity        uint `schema:"qty"`
}

// UpdateCart update shopping cart
func (ctrl Controller) UpdateCart(w http.ResponseWriter, req *http.Request) {
	var (
		input updateCartInput
		tx    = utils.GetDB(req)
	)

	req.ParseForm()
	decoder.Decode(&input, req.PostForm)

	order := getCurrentOrder(w, req)

	tx.Where(&orders.OrderItem{OrderID: order.ID, SizeVariationID: input.SizeVariationID}).
		Assign(&orders.OrderItem{Quantity: input.Quantity}).
		FirstOrCreate(&orders.OrderItem{})

	responder.With("html", func() {
		http.Redirect(w, req, "/cart", http.StatusFound)
	}).With([]string{"json", "xml"}, func() {
		config.Render.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
	}).Respond(req)
}

// Checkout checkout shopping cart
func (ctrl Controller) Checkout(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("checkout", map[string]interface{}{}, req, w)
}

// Complete complete shopping cart
func (ctrl Controller) Complete(w http.ResponseWriter, req *http.Request) {
}

// CheckoutSuccess checkout success page
func (ctrl Controller) CheckoutSuccess(w http.ResponseWriter, req *http.Request) {
	ctrl.View.Execute("success", map[string]interface{}{}, req, w)
}

func getCurrentOrder(w http.ResponseWriter, req *http.Request) *orders.Order {
	var (
		order       = orders.Order{}
		cardID      = manager.SessionManager.Get(req, "cart_id")
		currentUser = utils.GetCurrentUser(req)
		tx          = utils.GetDB(req)
	)

	if currentUser != nil {
		tx.First(&order, "id = ? AND user_id = ? OR user_id IS NULL", cardID, currentUser.ID)
		if order.UserID == 0 {
			tx.Model(&order).Update("UserID", currentUser.ID)
		}
	} else {
		tx.First(&order, "id = ? AND user_id IS NULL", cardID)
	}

	if order.State != orders.DraftState {
		order = orders.Order{}
		if currentUser != nil {
			order.UserID = currentUser.ID
		}

		tx.Create(&order)
	}

	manager.SessionManager.Add(w, req, "cart_id", order.ID)

	return &order
}
