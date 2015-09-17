package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Newsletter struct {
	gorm.Model
	UserID         uint
	User           User
	Email          string
	NewsletterType string
	MailType       string
	SubscribedAt   *time.Time
	UnsubscribedAt *time.Time
}
