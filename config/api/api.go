package api

import (
	"github.com/qor/admin"
	"github.com/qor/qor"

	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
)

var API *admin.Admin

func init() {
	API = admin.New(&qor.Config{DB: db.DB})

	Product := API.AddResource(&models.Product{})

	ColorVariationMeta := Product.Meta(&admin.Meta{Name: "ColorVariations"})
	ColorVariation := ColorVariationMeta.Resource
	ColorVariation.IndexAttrs("ID", "Color", "Images", "SizeVariations")
	ColorVariation.ShowAttrs("Color", "Images", "SizeVariations")

	SizeVariationMeta := ColorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	SizeVariation := SizeVariationMeta.Resource
	SizeVariation.IndexAttrs("ID", "Size", "AvailableQuantity")
	SizeVariation.ShowAttrs("ID", "Size", "AvailableQuantity")

	API.AddResource(&models.Order{})
	API.AddResource(&models.User{})
	API.AddResource(&models.Category{})
}
