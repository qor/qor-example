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
	OrderItems        []OrderItem
	transition.Transition
}

type OrderItem struct {
	gorm.Model
	OrderID         uint
	SizeVariationID uint
	SizeVariation   SizeVariation
	Quantity        uint
	Price           float32
	DiscountRate    uint
	transition.Transition
}

func (order Order) Amount() (amount float32) {
	for _, orderItem := range order.OrderItems {
		amount += orderItem.Amount()
	}
	return
}

func (orderItem OrderItem) Amount() float32 {
	return orderItem.Price * float32((100-orderItem.DiscountRate)*orderItem.Quantity)
}
