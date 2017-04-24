package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email                  string
	Password               string
	Name                   string
	Gender                 string
	Role                   string
	Balance                float32
	DefaultBillingAddress  uint
	DefaultShippingAddress uint
	Addresses              []Address

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time

	// Accepts
	AcceptPrivate bool
	AcceptLicense bool
	AcceptNews    bool
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN"}
}
