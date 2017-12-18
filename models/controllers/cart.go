package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/qor/qor"

	"github.com/qor/notification"
	"github.com/qor/notification/channels/database"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/cart"
	"github.com/qor/qor-example/config/utils"
)

// AddToCartHandler function
func AddToCartHandler(w http.ResponseWriter, req *http.Request) {
	var (
		curCart, _ = cart.GetCart(w, req)
		cartItem   cart.CartItem
	)
	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	if id, err := strconv.Atoi(req.FormValue("id")); err == nil {
		cartItem.ProductID = uint(id)
	} else {
		panic(err)
	}

	if qty, err := strconv.Atoi(req.FormValue("qty")); err == nil {
		cartItem.Quantity = uint(qty)
	} else {
		panic(err)
	}

	if _, err := curCart.Add(&cartItem); err != nil {
		panic(err)
	}

	utils.AddFlashMessage(w, req, fmt.Sprintf("Item added to cart."), "flash_success")

	http.Redirect(w, req, req.Referer(), http.StatusFound)
}

// EditCartHandler function
func EditCartHandler(w http.ResponseWriter, req *http.Request) {
	var (
		curCart, _ = cart.GetCart(w, req)
		id, _      = strconv.Atoi(utils.URLParam("id", req))
		action     = utils.URLParam("action", req)
		curItem    = curCart.GetContent()[uint(id)]
	)

	switch action {
	case "decrase":
		if (curItem.Quantity - 1) > 0 {
			curItem.Quantity = curItem.Quantity - 1
		}
	case "incrase":
		curItem.Quantity = curItem.Quantity + 1
	default:
		panic("unrecognized option")
	}

	if err := curCart.Edit(uint(id), curItem); err != nil {
		utils.AddFlashMessage(w, req, fmt.Sprintf("Can't edit item"), "flash_success")
	}

	http.Redirect(w, req, req.Referer(), http.StatusFound)
}

// RemoveFromCartHandler function
func RemoveFromCartHandler(w http.ResponseWriter, req *http.Request) {
	var (
		curCart, _ = cart.GetCart(w, req)
		id, _      = strconv.Atoi(utils.URLParam("id", req))
	)

	if err := curCart.Remove(uint(id)); err != nil {
		utils.AddFlashMessage(w, req, fmt.Sprintf("Can't delete item from cart."), "flash_error")
	} else {
		utils.AddFlashMessage(w, req, fmt.Sprintf("Item deleted from cart."), "flash_success")
	}

	http.Redirect(w, req, req.Referer(), http.StatusFound)
}

// ShowCartHandler function
func ShowCartHandler(w http.ResponseWriter, req *http.Request) {
	config.View.Execute("/cart/show", map[string]interface{}{}, req, w)
}

// CheckoutCartHandler function
func CheckoutCartHandler(w http.ResponseWriter, req *http.Request) {
	if curCart, _ := cart.GetCart(w, req); curCart.IsEmpty() {
		http.Redirect(w, req, "/cart", http.StatusTemporaryRedirect)
		return
	}

	var (
		address         models.Address
		deliveryMethods []models.DeliveryMethod
	)

	if utils.GetDB(req).Find(&deliveryMethods).RecordNotFound() {
		panic("Cannont find any delivery method")
	}

	if curUser := utils.GetCurrentUser(req); curUser != nil {
		utils.GetDB(req).Model(curUser).Related(&address)
	}

	config.View.Execute("/cart/checkout", map[string]interface{}{
		"DeliveryMethods": deliveryMethods,
		"Address":         address,
	}, req, w)
}

// OrderCartHandler function
func OrderCartHandler(w http.ResponseWriter, req *http.Request) {
	var (
		notifiTargets []models.User

		notify = notification.New(&notification.Config{})
		tx     = utils.GetDB(req)
	)

	if err := tx.Where(&models.User{Role: "Admin"}).Find(&notifiTargets).Error; err != nil {
		panic(err)
	}

	if err := req.ParseForm(); err != nil {
		panic(err)
	}

	var (
		deliveryMethodID uint
		paymentMethod    models.PaymentMethod
	)

	if val, err := strconv.Atoi(req.FormValue("delivery-method")); err != nil {
		panic(err)
	} else {
		deliveryMethodID = uint(val)

	}

	if val, err := strconv.Atoi(req.FormValue("payment_method")); err != nil {
		panic(err)
	} else {
		paymentMethod = models.PaymentMethod(uint8(val))
	}

	order := &models.Order{
		ShippingAddress: models.Address{
			Phone:    req.FormValue("address.phone"),
			City:     req.FormValue("address.city"),
			Address1: req.FormValue("address.address1"),
			Address2: req.FormValue("address.address2"),
		},
		DeliveryMethodID: deliveryMethodID,
		PaymentMethod:    paymentMethod,
	}

	tx.Model(&order).Related(&order.DeliveryMethod, "DeliveryMethod")

	if user := utils.GetCurrentUser(req); user != nil {
		order.UserID = user.ID
	}

	if curCart, err := cart.GetCart(w, req); err != nil {
		panic(err)
	} else {
		if curCart.IsEmpty() {
			panic(errors.New("Cart is empty"))
		}

		var svs = models.SizeVariations()

		tx.Where(curCart.GetItemsIDS()).Find(&svs)

		for _, sv := range svs {
			order.OrderItems = append(order.OrderItems, models.OrderItem{
				SizeVariation: sv,
				Quantity:      curCart.GetContent()[sv.ID].Quantity,
				Price:         sv.ColorVariation.Product.Price,
			})
		}

		order.PaymentAmount = order.Amount()
		order.PaymentTotal = order.Total()

		if err := curCart.EmptyCart(); err != nil {
			panic(err)
		}
	}

	if err := tx.Create(&order).Error; err != nil {
		panic(err)
	}

	if err := models.OrderState.Trigger("checkout", order, tx); err != nil {
		panic(err)
	}

	if err := tx.Save(&order).Error; err != nil {
		panic(err)
	}

	notify.RegisterChannel(database.New(&database.Config{DB: tx}))
	for _, target := range notifiTargets {
		notify.Send(&notification.Message{
			From:        utils.GetCurrentUser(req),
			To:          target,
			Title:       "New order",
			Body:        fmt.Sprintf("Created order #%d, total: $%f", order.ID, order.PaymentTotal),
			MessageType: "order_created",
		}, &qor.Context{DB: tx})
	}

	utils.AddFlashMessage(w, req, fmt.Sprintf("Order created"), "flash_success")

	if user := utils.GetCurrentUser(req); user != nil {
		http.Redirect(w, req, "/account", http.StatusFound)
	} else {
		http.Redirect(w, req, "/", http.StatusFound)
	}
}
