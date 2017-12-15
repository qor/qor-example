package api

import (
	"github.com/qor/admin"
	"github.com/qor/qor"

	"github.com/qor/qor-example/app/orders"
	"github.com/qor/qor-example/app/products"
	"github.com/qor/qor-example/app/users"
	"github.com/qor/qor-example/db"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})

	Product := API.AddResource(&products.Product{})

	ColorVariationMeta := Product.Meta(&admin.Meta{Name: "ColorVariations"})
	ColorVariation := ColorVariationMeta.Resource
	ColorVariation.IndexAttrs("ID", "Color", "Images", "SizeVariations")
	ColorVariation.ShowAttrs("Color", "Images", "SizeVariations")

	SizeVariationMeta := ColorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	SizeVariation := SizeVariationMeta.Resource
	SizeVariation.IndexAttrs("ID", "Size", "AvailableQuantity")
	SizeVariation.ShowAttrs("ID", "Size", "AvailableQuantity")

	API.AddResource(&orders.Order{})

	User := API.AddResource(&users.User{})
	userOrders, _ := User.AddSubResource("Orders")
	userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})

	API.AddResource(&products.Category{})
}
