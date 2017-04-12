package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/qor/action_bar"
	// "github.com/qor/qor"
	"dukeondope.ru/mlm/sandbox/app/models"
	"dukeondope.ru/mlm/sandbox/config"
	// "dukeondope.ru/mlm/sandbox/config/admin"
	// "dukeondope.ru/mlm/sandbox/config/seo"
)

func CabinetShow(ctx *gin.Context) {
	var (
		currentUser = CurrentUser(ctx)
		addresses   []models.Address
	)

	DB(ctx).Where(models.Address{UserID: currentUser.ID}).Find(&addresses)

	fmt.Println(addresses)

	config.View.Funcs(funcsMap(ctx)).Execute(
		"cabinet_show",
		gin.H{
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

	fmt.Println(billingAddress)

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

	fmt.Println(shippingAddress)

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"status":  http.StatusCreated,
			"message": "Shipping address added successfully!",
		},
	)
}
