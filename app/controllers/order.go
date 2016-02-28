package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	// "github.com/gin-gonic/contrib/sessions"
)

func removeDuplicates(a []uint) []uint {
	result := []uint{}
	seen := map[uint]uint{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			result = append(result, val)
			seen[val] = val
		}
	}
	return result
}

func OrderIndex(ctx *gin.Context) {
	var orders []models.Order
	// session := sessions.Default(ctx)

	if err := db.DB.Preload("Store.User").Preload("Car.Drivers").Preload("OrderItems").Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Bad request"})

	} else {

		// var invoices []models.Invoice
		// var invoice models.Invoice
		// var cardsId []uint
		// for _, order := range orders {
		// 	paymentAmount := order.PaymentAmount
		// 	var items []models.InvoiceList
		// 	var rfids []models.Rfid
		// 	var uids []uint
		// 	var rfidNames []string

		// 	for _, item := range order.OrderItems {
		// 		amount := item.Quantity * item.Price
		// 		line := models.InvoiceList{
		// 			Name:     item.Name,
		// 			Quantity: item.Quantity,
		// 			Price:    item.Price,
		// 			Amount:   amount,
		// 		}
		// 		paymentAmount = paymentAmount - amount
		// 		items = append(items, line)
		// 	}
		// 	for _, uid := range order.Car.Drivers {
		// 		uids = append(uids, uid.ID)
		// 		cardsId = append(cardsId, uid.ID)
		// 	}
		// 	for _, uid := range order.Store.User {
		// 		uids = append(uids, uid.ID)
		// 		cardsId = append(cardsId, uid.ID)
		// 	}
		// 	if err := db.DB.Where("user_id IN (?)", uids).Find(&rfids).Error; err == nil {
		// 		// fmt.Println("card ID", uids)
		// 		for _, item := range rfids {
		// 			// fmt.Println(rfids)
		// 			rfidNames = append(rfidNames, item.Name)
		// 		}
		// 	}
		// 	invoice = models.Invoice{
		// 		ID:         order.ID,
		// 		Name:       order.Name,
		// 		AutoNumber: order.Car.CarNumber,
		// 		Partner:    order.Store.Name,
		// 		Cards:      rfidNames,
		// 		Barcode:    false,
		// 		Amount:     order.PaymentAmount,
		// 		Invoices:   items,
		// 	}
		// 	invoices = append(invoices, invoice)
		// 	if paymentAmount > float32(0) {
		// 		fmt.Println("[ERROR] PaymentAmount", paymentAmount, ">0")
		// 	}
		// }
		// cards := removeDuplicates(cardsId)
		// fmt.Println(removeDuplicates(cardsId))
		// fmt.Printf("Client ip: %s\n", ctx.ClientIP())
		// ctx.JSON(http.StatusOK, gin.H{"status": "success", "invoices": &invoices, "cards": &cards})
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": &orders})
	}

	// ctx.HTML(
	// 	http.StatusOK,
	// 	"product_index.tmpl",
	// 	gin.H{
	// 		"orders": orders,
	// 		"seoTag":   seoObj.DefaultPage.Render(seoObj, nil),
	// 		"microSearch": seo.MicroSearch{
	// 			URL:    "http://demo.getqor.com",
	// 			Target: "http://demo.getqor.com/search?q=",
	// 		}.Render(),
	// 		"microContact": seo.MicroContact{
	// 			URL:         "http://demo.getqor.com",
	// 			Telephone:   "080-0012-3232",
	// 			ContactType: "Customer Service",
	// 		}.Render(),
	// 	},
	// )
}

// func OrderShow(ctx *gin.Context) {
// 	var order models.Order
// 	db.DB.Preload("ColorVariations").Preload("ColorVariations.Images").Find(&order, ctx.Param("id"))
// 	seoObj := models.Seo{}
// 	db.DB.First(&seoObj)

// 	var imageURL string
// 	if len(order.ColorVariations) > 0 && len(order.ColorVariations[0].Images) > 0 {
// 		imageURL = order.ColorVariations[0].Images[0].Image.URL()
// 	}

// 	ctx.HTML(
// 		http.StatusOK,
// 		"order_show.tmpl",
// 		gin.H{
// 			"order": order,
// 			"seoTag":  seoObj.ProductPage.Render(seoObj, order),
// 			"microProduct": seo.MicroProduct{
// 				Name:        product.Name,
// 				Description: product.Description,
// 				BrandName:   product.Category.Name,
// 				SKU:         product.Code,
// 				Price:       float64(product.Price),
// 				Image:       imageURL,
// 			}.Render(),
// 		},
// 	)
// }
