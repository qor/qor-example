package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type AuthIdentity struct {
	gorm.Model
	Provider          string // phone, email, wechat, github...
	UID               string
	EncryptedPassword string
	AuthInfo          AuthInfo
	UserID            string
	State             string // unconfirmed, confirmed, expired

	Password             string `gorm:"-"`
	PasswordConfirmation string `gorm:"-"`
}

////////////////////////////////////////////////////////////////////////////////
type SignLog struct {
	UserAgent string
	At        *time.Time
	IP        string
}

type AuthInfo struct {
	PhoneVerificationCode       string
	PhoneVerificationCodeExpiry *time.Time
	PhoneConfirmedAt            *time.Time
	UnconfirmedPhone            string // only use when changing phone number

	EmailConfirmedAt *time.Time
	UnconfirmedEmail string // only use when changing email

	SignInCount uint
	SignLogs    []SignLog
}
