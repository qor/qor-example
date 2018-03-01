package orders

import (
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor-example/models/products"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/transition"
)

type PaymentMethod uint8

const (
	COD = PaymentMethod(iota)
)

func (pm PaymentMethod) String() string {
	name := []string{"COD"}
	i := uint8(pm)
	switch {
	case i < uint8(len(name)):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

type Order struct {
	gorm.Model
	UserID            uint
	User              users.User
	PaymentAmount     float32
	PaymentTotal      float32
	AbandonedReason   string
	DiscountValue     uint
	DeliveryMethodID  uint `form:"delivery-method"`
	DeliveryMethod    DeliveryMethod
	PaymentMethod     PaymentMethod
	TrackingNumber    *string
	ShippedAt         *time.Time
	ReturnedAt        *time.Time
	CancelledAt       *time.Time
	ShippingAddressID uint `form:"shippingaddress"`
	ShippingAddress   users.Address
	BillingAddressID  uint `form:"billingaddress"`
	BillingAddress    users.Address
	OrderItems        []OrderItem
	transition.Transition
}

type OrderItem struct {
	gorm.Model
	OrderID         uint
	SizeVariationID uint `cartitem:"SizeVariationID"`
	SizeVariation   products.SizeVariation
	Quantity        uint `cartitem:"Quantity"`
	Price           float32
	DiscountRate    uint
	transition.Transition
}

func (order Order) IsCart() bool {
	return order.State == DraftState || order.State == ""
}

func (order Order) Amount() (amount float32) {
	for _, orderItem := range order.OrderItems {
		amount += orderItem.Amount()
	}
	return
}

func (order Order) Total() (total float32) {
	total = order.Amount() - float32(order.DiscountValue)
	total = order.Amount() + float32(order.DeliveryMethod.Price)
	return
}

func (item OrderItem) Amount() float32 {
	return item.Price * float32(item.Quantity) * float32(100-item.DiscountRate) / 100
}
