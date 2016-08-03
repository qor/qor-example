package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Password  string
	Name      sql.NullString
	Gender    string
	Role      string
	Addresses []Address

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "zh-CN"}
}
