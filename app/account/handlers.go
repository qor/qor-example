package account

import (
	"net/http"

	"github.com/qor/qor-example/models/orders"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/qor-example/utils"
	"github.com/qor/render"
)

// Controller products controller
type Controller struct {
	View *render.Render
}

// Profile profile show page
func (ctrl Controller) Profile(w http.ResponseWriter, req *http.Request) {
	var (
		currentUser                     = utils.GetCurrentUser(req)
		tx                              = utils.GetDB(req)
		billingAddress, shippingAddress users.Address
	)

	// TODO refactor
	tx.Model(currentUser).Related(&currentUser.Addresses, "Addresses")
	tx.Model(currentUser).Related(&billingAddress, "DefaultBillingAddress")
	tx.Model(currentUser).Related(&shippingAddress, "DefaultShippingAddress")

	ctrl.View.Execute("profile", map[string]interface{}{
		"CurrentUser": currentUser, "DefaultBillingAddress": billingAddress, "DefaultShippingAddress": shippingAddress,
	}, req, w)
}

// Orders orders page
func (ctrl Controller) Orders(w http.ResponseWriter, req *http.Request) {
	var (
		Orders      []orders.Order
		currentUser = utils.GetCurrentUser(req)
		tx          = utils.GetDB(req)
	)

	tx.Preload("OrderItems").Where("state <> ? AND state != ?", orders.DraftState, "").Where(&orders.Order{UserID: &currentUser.ID}).Find(&Orders)

	ctrl.View.Execute("orders", map[string]interface{}{"Orders": Orders}, req, w)
}

// Update update profile page
func (ctrl Controller) Update(w http.ResponseWriter, req *http.Request) {
	// FIXME
}

// AddCredit add credit
func (ctrl Controller) AddCredit(w http.ResponseWriter, req *http.Request) {
	// FIXME
}
