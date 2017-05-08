package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	// "github.com/qor/action_bar"
	// "github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	// "github.com/qor/qor-example/config/admin"
	// "github.com/qor/qor-example/config/seo"
)

func CabinetShow(ctx *gin.Context) {
	var (
		currentUser = CurrentUser(ctx)
		orders      []models.Order
		session     = sessions.Default(ctx)
		flashes     = session.Flashes()
	)
	session.Save()

	DB(ctx).Preload("OrderItems").Preload("OrderItems.SizeVariation.Size").Preload("OrderItems.SizeVariation.ColorVariation.Color").Preload("OrderItems.SizeVariation.ColorVariation.Product").Where(&models.Order{UserID: currentUser.ID}).Find(&orders)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cabinet/cabinet_show",
		gin.H{
			"Flashes":       flashes,
			"Orders":        orders,
			"CurrentUser":   currentUser,
			"CurrentLocale": CurrentLocale(ctx),
			"Categories":    CategoriesList(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func AddUserCredit(ctx *gin.Context) {
	var (
		user    = CurrentUser(ctx)
		session = sessions.Default(ctx)
	)

	fmt.Println(user)

	session.AddFlash("Cash added successfully!")
	session.Save()

	http.Redirect(ctx.Writer, ctx.Request, "/cabinet", http.StatusFound)
}

func ProfileShow(ctx *gin.Context) {
	var (
		user            models.User
		billingAddress  models.Address
		shippingAddress models.Address
		session         = sessions.Default(ctx)
		flashes         = session.Flashes()
	)

	session.Save()

	DB(ctx).Preload("Addresses").Where(CurrentUser(ctx)).First(&user)

	DB(ctx).First(&billingAddress, user.DefaultBillingAddress)
	DB(ctx).First(&shippingAddress, user.DefaultShippingAddress)

	// fmt.Println(session.Flashes())
	// session.Save()

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cabinet/profile_show",
		gin.H{
			"Flashes":                flashes,
			"CurrentUser":            user,
			"CurrentLocale":          CurrentLocale(ctx),
			"Categories":             CategoriesList(ctx),
			"DefaultBillingAddress":  billingAddress,
			"DefaultShippingAddress": shippingAddress,
		},
		ctx.Request,
		ctx.Writer,
	)
}

func SetUserProfile(ctx *gin.Context) {
	var (
		user    models.User
		session = sessions.Default(ctx)
	)

	DB(ctx).Preload("Addresses").Where(CurrentUser(ctx).ID).First(&user)

	// 	!!!! REMOVE THIS
	user.AcceptPrivate = false
	user.AcceptLicense = false
	user.AcceptNews = false
	ctx.Bind(&user)

	fmt.Println(user)
	DB(ctx).Save(&user)

	session.AddFlash("Profile updated successfully!")
	session.Save()

	http.Redirect(ctx.Writer, ctx.Request, "/cabinet/profile", http.StatusFound)
	/* 	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Profile updated successfully!",
		},
	) */
}

func SetBillingAddress(ctx *gin.Context) {
	var (
		billingAddress models.Address
		user           = CurrentUser(ctx)
		session        = sessions.Default(ctx)
	)

	DB(ctx).First(&billingAddress, user.DefaultBillingAddress)
	ctx.Bind(&billingAddress)
	DB(ctx).Save(&billingAddress)

	session.AddFlash("Address updated successfully!")
	session.Save()

	http.Redirect(ctx.Writer, ctx.Request, "/cabinet/profile", http.StatusFound)
	/* 	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Billing address added successfully!",
		},
	) */
}

func SetShippingAddress(ctx *gin.Context) {
	var (
		shippingAddress models.Address
		user            = CurrentUser(ctx)
		session         = sessions.Default(ctx)
	)

	DB(ctx).First(&shippingAddress, user.DefaultShippingAddress)
	ctx.Bind(&shippingAddress)
	DB(ctx).Save(&shippingAddress)

	session.AddFlash("Address updated successfully!")
	session.Save()

	http.Redirect(ctx.Writer, ctx.Request, "/cabinet/profile", http.StatusFound)

	/* 	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Shipping address added successfully!",
		},
	) */
}
