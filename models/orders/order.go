package orders

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/qor/qor-example/models/users"
	"github.com/qor/transition"
)

type PaymentMethod = string

const (
	COD        PaymentMethod = "COD"
	AmazonPay  PaymentMethod = "AmazonPay"
	CreditCard PaymentMethod = "CreditCard"
)

type Order struct {
	gorm.Model
	UserID                   *uint
	User                     users.User
	PaymentAmount            float32
	PaymentTotal             float32
	AbandonedReason          string
	DiscountValue            uint
	DeliveryMethodID         uint `form:"delivery-method"`
	DeliveryMethod           DeliveryMethod
	PaymentMethod            string
	TrackingNumber           *string
	ShippedAt                *time.Time
	ReturnedAt               *time.Time
	CancelledAt              *time.Time
	ShippingAddressID        uint `form:"shippingaddress"`
	ShippingAddress          users.Address
	BillingAddressID         uint `form:"billingaddress"`
	BillingAddress           users.Address
	OrderItems               []OrderItem
	AmazonAddressAccessToken string
	AmazonOrderReferenceID   string
	AmazonAuthorizationID    string
	AmazonCaptureID          string
	AmazonRefundID           string
	PaymentLog               string `gorm:"size:655250"`
	transition.Transition
}

func (order Order) ExternalID() string {
	return fmt.Sprint(order.ID)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (order Order) UniqueExternalID() string {
	return fmt.Sprint(order.ID) + "_" + randomString(6)
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

// DeliveryFee delivery fee
func (order Order) DeliveryFee() (amount float32) {
	return order.DeliveryMethod.Price
}

func (order Order) Total() (total float32) {
	total = order.Amount() - float32(order.DiscountValue)
	total = order.Amount() + float32(order.DeliveryMethod.Price)
	return
}
