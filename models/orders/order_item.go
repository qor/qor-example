package orders

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor-example/config/db"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/transition"
)

type OrderItem struct {
	gorm.Model
	OrderID         uint
	SizeVariationID uint `cartitem:"SizeVariationID"`
	SizeVariation   *products.SizeVariation
	Quantity        uint `cartitem:"Quantity"`
	Price           float32
	DiscountRate    uint
	transition.Transition
}

// IsCart order item's state is cart
func (item OrderItem) IsCart() bool {
	return item.State == DraftState || item.State == ""
}

func (item *OrderItem) loadSizeVariation() {
	if item.SizeVariation == nil {
		item.SizeVariation = &products.SizeVariation{}
		db.DB.Model(item).Preload("Size").Preload("ColorVariation.Product").Preload("ColorVariation.Color").Association("SizeVariation").Find(item.SizeVariation)
	}
}

// ProductImageURL get product image
func (item *OrderItem) ProductImageURL() string {
	item.loadSizeVariation()
	return item.SizeVariation.ColorVariation.MainImageURL()
}

// SellingPrice order item's selling price
func (item *OrderItem) SellingPrice() float32 {
	if item.IsCart() {
		item.loadSizeVariation()
		return item.SizeVariation.ColorVariation.Product.Price
	}
	return item.Price
}

// ProductName order item's color name
func (item *OrderItem) ProductName() string {
	item.loadSizeVariation()
	return item.SizeVariation.ColorVariation.Product.Name
}

// ColorName order item's color name
func (item *OrderItem) ColorName() string {
	item.loadSizeVariation()
	return item.SizeVariation.ColorVariation.Color.Name
}

// SizeName order item's size name
func (item *OrderItem) SizeName() string {
	item.loadSizeVariation()
	return item.SizeVariation.Size.Name
}

// ProductPath order item's product name
func (item *OrderItem) ProductPath() string {
	item.loadSizeVariation()
	return item.SizeVariation.ColorVariation.ViewPath()
}

// Amount order item's amount
func (item OrderItem) Amount() float32 {
	amount := item.SellingPrice() * float32(item.Quantity)
	if item.DiscountRate > 0 && item.DiscountRate <= 100 {
		amount = amount * float32(100-item.DiscountRate) / 100
	}
	return amount
}
