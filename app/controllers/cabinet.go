package controllers

import (
	"net/http"

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
		addresses   []models.Address
	)

	DB(ctx).Model(&currentUser).Related(&addresses)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cabinet/cabinet_show",
		gin.H{
			"CurrentUser":   currentUser,
			"CurrentLocale": CurrentLocale(ctx),
			"Categories":    CategoriesList(ctx),
			"Addresses":     addresses,
		},
		ctx.Request,
		ctx.Writer,
	)
}

func ProfileShow(ctx *gin.Context) {
	var (
		currentUser = CurrentUser(ctx)
		orders      []models.Order
	)

	DB(ctx).Preload("OrderItems").Preload("OrderItems.SizeVariation.Size").Preload("OrderItems.SizeVariation.ColorVariation.Color").Where(&models.Order{UserID: currentUser.ID}).Find(&orders)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cabinet/profile_show",
		gin.H{
			"Orders":        orders,
			"CurrentUser":   currentUser,
			"CurrentLocale": CurrentLocale(ctx),
			"Categories":    CategoriesList(ctx),
		},
		ctx.Request,
		ctx.Writer,
	)
}

func SetBillingAddress(ctx *gin.Context) {
	var (
		billingAddress models.Address
	)

	ctx.Bind(&billingAddress)
	billingAddress.UserID = CurrentUser(ctx).ID

	DB(ctx).Create(&billingAddress)

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Billing address added successfully!",
		},
	)
}

func SetShippingAddress(ctx *gin.Context) {
	var (
		shippingAddress models.Address
	)

	ctx.Bind(&shippingAddress)
	shippingAddress.UserID = CurrentUser(ctx).ID

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Shipping address added successfully!",
		},
	)
}
