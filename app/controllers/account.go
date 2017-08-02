package controllers

import (
	"fmt"
	"net/http"

	// "github.com/qor/action_bar"
	// "github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/config/utils"
	// "github.com/qor/qor-example/config/admin"
	// "github.com/qor/qor-example/config/seo"
)

func AccountShow(w http.ResponseWriter, req *http.Request) {
	var (
		orders      []models.Order
		currentUser = utils.GetCurrentUser(req)
		tx          = utils.GetDB(req)
	)

	tx.Preload("OrderItems").Preload("OrderItems.SizeVariation.Size").Preload("OrderItems.SizeVariation.ColorVariation.Color").Preload("OrderItems.SizeVariation.ColorVariation.Product").Where(&models.Order{UserID: currentUser.ID}).Find(&orders)

	config.View.Execute(
		"account/show",
		map[string]interface{}{"Orders": orders},
		req,
		w,
	)
}

func AddUserCredit(w http.ResponseWriter, req *http.Request) {
	// TODO
	// session.AddFlash("Cash added successfully!")

	http.Redirect(w, req, "/account", http.StatusFound)
}

func ProfileShow(w http.ResponseWriter, req *http.Request) {
	var (
		currentUser                     = utils.GetCurrentUser(req)
		tx                              = utils.GetDB(req)
		billingAddress, shippingAddress models.Address
	)

	// TODO refactor
	tx.Model(currentUser).Related(&currentUser.Addresses, "Addresses")
	tx.Model(currentUser).Related(&billingAddress, "DefaultBillingAddress")
	tx.Model(currentUser).Related(&shippingAddress, "DefaultShippingAddress")

	config.View.Execute(
		"account/profile_show",
		map[string]interface{}{
			"CurrentUser":            currentUser,
			"DefaultBillingAddress":  billingAddress,
			"DefaultShippingAddress": shippingAddress,
		},
		req,
		w,
	)
}

func SetUserProfile(w http.ResponseWriter, req *http.Request) {
	var (
		user = utils.GetCurrentUser(req)
		tx   = utils.GetDB(req)
	)

	tx.Model(user).Related(&user.Addresses, "Addresses")

	// 	!!!! REMOVE THIS
	user.AcceptPrivate = false
	user.AcceptLicense = false
	user.AcceptNews = false

	// FIXME DECODE ctx.Bind(&user)

	// TODO
	// session.AddFlash("Profile updated successfully!")

	fmt.Println(user)
	tx.Save(&user)

	http.Redirect(w, req, "/account/profile", http.StatusFound)
}

func SetBillingAddress(w http.ResponseWriter, req *http.Request) {
	var (
		billingAddress models.Address
		// user           = utils.GetCurrentUser(req)
		tx = utils.GetDB(req)
	)

	// FIXME
	// ctx.Bind(&billingAddress)
	tx.Save(&billingAddress)

	// TODO
	// session.AddFlash("Address updated successfully!")
	// session.Save()

	http.Redirect(w, req, "/account/profile", http.StatusFound)
}

func SetShippingAddress(w http.ResponseWriter, req *http.Request) {
	var (
		shippingAddress models.Address
		// user            = utils.GetCurrentUser(req)
		tx = utils.GetDB(req)
	)

	// FIXME
	// ctx.Bind(&shippingAddress)
	tx.Save(&shippingAddress)

	// TODO
	// session.AddFlash("Address updated successfully!")
	// session.Save()

	http.Redirect(w, req, "/account/profile", http.StatusFound)
}
