package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/qor/transition"
)

type Order struct {
	gorm.Model
	UserID            uint
	DiscountValue     uint
	ShippingAddressID uint
	ShippingAddress   Address
	BillingAddressID  uint
	BillingAddress    Address
	transition.Transition
}

type OrderItem struct {
	gorm.Model
	SizeVariationID uint
	SizeVariation   SizeVariation
	Quantity        uint
	Price           float32
	DiscountRate    uint
	transition.Transition
}
