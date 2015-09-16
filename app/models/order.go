package models

import (
	"errors"

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

func (item OrderItem) Amount() float32 {
	return item.Price * float32(item.Quantity) * float32(100-item.DiscountRate) / 100
}

var OrderStateMachine = transition.New(&OrderItem{})

func init() {
	OrderStateMachine.Initial("draft")
	OrderStateMachine.State("checkout")
	OrderStateMachine.State("paid").Enter(func(value interface{}, tx *gorm.DB) error {
		// freeze stock
		return nil
	})
	OrderStateMachine.State("cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
		// release stock
		return nil
	})
	OrderStateMachine.State("paid_cancelled").Enter(func(value interface{}, tx *gorm.DB) error {
		if _, ok := value.(*OrderItem); ok {
			// do refund
			return nil
		}
		return errors.New("not order item")
	})
	OrderStateMachine.State("processing")
	OrderStateMachine.State("shipped").Enter(func(value interface{}, tx *gorm.DB) error {
		// send shipment email
		return nil
	})
	OrderStateMachine.State("returned").Enter(func(value interface{}, tx *gorm.DB) error {
		// do refund
		return nil
	})

	OrderStateMachine.Event("checkout").To("checkout").From("draft")
	OrderStateMachine.Event("pay").To("paid").From("checkout")
	cancelEvent := OrderStateMachine.Event("cancel")
	cancelEvent.To("cancelled").From("checkout")
	cancelEvent.To("paid_cacelled").From("paid")

	OrderStateMachine.Event("process").To("processing").From("paid")
	OrderStateMachine.Event("ship").To("shipped").From("processing")
	OrderStateMachine.Event("return").To("returned").From("shipped")
}
