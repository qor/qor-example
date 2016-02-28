package controllers

import (
	"net/http"

	"github.com/apertoire/mlog"
	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
	// "github.com/gin-gonic/contrib/sessions"
)

// GET: http://localhost:7000/api/v1/category
func CategoryIndex(ctx *gin.Context) {
	mlog.Start(mlog.LevelTrace, "")
	var categorys []models.Category
	// ctx.Header("Locale", value)
	acceptLanguage := ctx.Request.Header.Get("Accept-Language")[0:2]
	locale := ctx.Request.Header.Get("Locale")
	if len(locale) == 0 {
		locale = config.Config.Locale
	}
	// ctx.Logger
	mlog.Trace("acceptLanguage: %v, locale: %v", acceptLanguage, locale)
	// d := db.DB.Set("l10n:locale", "ru-RU")

	// session := sessions.Default(ctx)

	// mode := "locale"
	// if err := d.Set("l10n:mode", mode).Find(&categorys).Error; err != nil {
	if err := db.DB.Set("l10n:locale", locale).Find(&categorys).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})
	} else {

		// 	var invoices []models.Invoice
		// 	var invoice models.Invoice
		// 	var cardsId []uint
		// 	for _, order := range orders {
		// 		paymentAmount := order.PaymentAmount
		// 		var items []models.InvoiceList
		// 		var rfids []models.Rfid
		// 		var uids []uint
		// 		var rfidNames []string

		// 		for _, item := range order.OrderItems {
		// 			amount := item.Quantity * item.Price
		// 			line := models.InvoiceList{
		// 				Name:     item.Name,
		// 				Quantity: item.Quantity,
		// 				Price:    item.Price,
		// 				Amount:   amount,
		// 			}
		// 			paymentAmount = paymentAmount - amount
		// 			items = append(items, line)
		// 		}
		// 		for _, uid := range order.Car.Drivers {
		// 			uids = append(uids, uid.ID)
		// 			cardsId = append(cardsId, uid.ID)
		// 		}
		// 		for _, uid := range order.Store.User {
		// 			uids = append(uids, uid.ID)
		// 			cardsId = append(cardsId, uid.ID)
		// 		}
		// 		if err := db.DB.Where("user_id IN (?)", uids).Find(&rfids).Error; err == nil {
		// 			// fmt.Println("card ID", uids)
		// 			for _, item := range rfids {
		// 				// fmt.Println(rfids)
		// 				rfidNames = append(rfidNames, item.Name)
		// 			}
		// 		}
		// 		invoice = models.Invoice{
		// 			ID:         order.ID,
		// 			Name:       order.Name,
		// 			AutoNumber: order.Car.CarNumber,
		// 			Partner:    order.Store.Name,
		// 			Cards:      rfidNames,
		// 			Barcode:    false,
		// 			Amount:     order.PaymentAmount,
		// 			Invoices:   items,
		// 		}
		// 		invoices = append(invoices, invoice)
		// 		if paymentAmount > float32(0) {
		// 			fmt.Println("[ERROR] PaymentAmount", paymentAmount, ">0")
		// 		}
	}
	// 	cards := removeDuplicates(cardsId)
	// 	// fmt.Println(removeDuplicates(cardsId))
	// 	fmt.Printf("Client ip: %s\n", ctx.ClientIP())
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &categorys})
}
